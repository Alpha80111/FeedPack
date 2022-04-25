package server

import (
	"encoding/json"
	"enterpret/dataaccess"
	"enterpret/models"
	"enterpret/sources"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct {
	client  http.Client
	logger  *log.Logger
	ds      dataaccess.DataStore
	sources sources.SourceProcessor
}

func (s *server) init() error {

	http.HandleFunc("/push/message", s.handleMessage)

	err := http.ListenAndServe("localhost:8088", nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) handleMessage(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method Not Allowed"))
		if err != nil {
			s.logger.Printf("Method %s Not Allowed", req.Method)
			return
		}
		return
	}

	p, err := io.ReadAll(req.Body)
	if err != nil {
		s.badRequestError(w, err)
		return
	}
	var reqBody models.PushRequestBody
	err = json.Unmarshal(p, &reqBody)
	if err != nil {
		s.badRequestError(w, err)
		return
	}

	if mp, err := s.sources.GetProcessor(reqBody.Source); err != nil {
		s.badRequestError(w, err)
		return
	} else {
		_, err := mp.IngestAndStoreFeedback(p, reqBody.Tenant)
		if err != nil {
			s.badRequestError(w, err)
			return
		}
	}

	w.Write([]byte(`Successfully processed message`))
}

func (s *server) badRequestError(w http.ResponseWriter, er error) {
	w.WriteHeader(http.StatusBadRequest)
	errStr := ""
	if er != nil {
		errStr = er.Error()
	}
	_, err := w.Write([]byte("Bad Request: " + errStr))
	if err != nil {
		s.logger.Printf("Invalid body received")
		return
	}
	return
}

func NewServer(store dataaccess.DataStore) error {
	s := server{
		//client:  http.Client{},
		logger:  log.New(os.Stdout, "logger: ", 1),
		ds:      store,
		sources: sources.NewSourceProcessor(store),
	}

	return s.init()
}
