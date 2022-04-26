package server

import (
	"encoding/json"
	"enterpret/dataaccess"
	"enterpret/models"
	"enterpret/sources"
	"io"
	"log"
	"net/http"
)

type server struct {
	logger  *log.Logger
	ds      dataaccess.DataStore
	sources sources.SourceProcessor
}

func (s *server) handleFeedback(w http.ResponseWriter, req *http.Request) {

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

	if fp, err := s.sources.GetProcessor(reqBody.Source); err != nil {
		s.badRequestError(w, err)
		return
	} else {
		_, err := fp.IngestAndStoreFeedback(p, reqBody.Tenant)
		if err != nil {
			s.badRequestError(w, err)
			return
		}
	}

	_, err = w.Write([]byte(`Successfully processed message`))
	if err != nil {
		s.logger.Println("Push Handler: Successfully processed and stored message but failed to send response for req: tenant ",
			reqBody.Tenant, " source: ", reqBody.Source)
		return
	}
	s.logger.Println("Push Handler: Successfully processed and stored message for req: tenant ",
		reqBody.Tenant, " source: ", reqBody.Source)
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
