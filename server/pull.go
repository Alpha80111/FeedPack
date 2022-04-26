package server

import (
	"encoding/json"
	"enterpret/models"
	"io"
	"log"
	"net/http"
)

func (s *server) handlePullRequest(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("Method Not Allowed"))
		if err != nil {
			s.logger.Printf("Method %s Not Allowed", req.Method)
			return
		}
		return
	}

	payload, err := io.ReadAll(req.Body)
	if err != nil {
		s.badRequestError(w, err)
		return
	}
	var reqBody models.PullRequestBody
	err = json.Unmarshal(payload, &reqBody)
	if err != nil {
		s.badRequestError(w, err)
		return
	}

	fp, err := s.sources.GetProcessor(reqBody.Source)
	if err != nil {
		log.Default().Println("Failed fetching message processor for source "+reqBody.Source, err.Error())
		return
	}

	_, err = fp.FetchAndStoreFeedbacks(reqBody.Params, reqBody.Tenant)
	if err != nil {
		log.Default().Println("Failed fetching and storing messages: ", err.Error())
		return
	}

	w.Write([]byte(`Successfully processed message`))
}
