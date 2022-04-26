package server

import (
	"enterpret/dataaccess"
	"enterpret/sources"
	"log"
	"net/http"
	"os"
)

func (s *server) init() error {

	http.HandleFunc("/push/feedback", s.handleFeedback)
	http.HandleFunc("/pull/feedback", s.handlePullRequest)
	http.HandleFunc("/fetch/feedbacks", s.handleFetchRequest)

	err := http.ListenAndServe("localhost:8088", nil)
	if err != nil {
		return err
	}

	return nil
}

func NewServer(store dataaccess.DataStore, processor sources.SourceProcessor) error {
	s := server{
		logger:  log.New(os.Stdout, "logger: ", 1),
		ds:      store,
		sources: processor,
	}

	return s.init()
}
