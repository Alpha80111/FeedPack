package server

import (
	"encoding/json"
	"enterpret/dataaccess"
	"enterpret/sources"
	"fmt"
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

func (s *server) init() {

	http.HandleFunc("/push/message", s.handleMessage)

	err := http.ListenAndServe("localhost:8088", nil)
	if err != nil {
		return
	}
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

	//for name, headers := range req.Header {
	//	for _, h := range headers {
	//		fmt.Printf("%v: %v\n", name, h)
	//	}
	//}

	var p []byte
	req.Body.Read(p)
	fmt.Println(string(p))
	var reqBody map[string]string
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Bad Request"))
		if err != nil {
			s.logger.Printf("Invalid body received")
			return
		}
		return
	}

	fmt.Println(reqBody)

	if reqBody["source"] == "discourse" {

	}
	w.Write([]byte(`Successfully received message`))
}

func NewServer(store dataaccess.DataStore) {
	s := server{
		//client:  http.Client{},
		logger:  log.New(os.Stdout, "logger: ", 1),
		ds:      store,
		sources: sources.NewSourceProcessor(store),
	}

	s.init()
}
