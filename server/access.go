package server

import (
	"encoding/json"
	"enterpret/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

func (s *server) handleFetchRequest(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
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
	var reqBody models.AccessRequestBody
	err = json.Unmarshal(payload, &reqBody)
	if err != nil {
		s.badRequestError(w, err)
		return
	}

	fmt.Println(reqBody)

	sort.Strings(reqBody.Sources)
	if reqBody.Records < 1 {
		reqBody.Records = 100
	}

	c, err := s.ds.FetchFeedbacks(reqBody.Tenant, reqBody.Page, reqBody.Records, reqBody.Sources...)
	for _, m := range c {
		log.Default().Println(m.Meta.ID)
	}
	respBody, err := json.Marshal(c)
	if err != nil {
		s.logger.Printf("Error marshalling feedback records")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Unable to fetch records due to internal server error`))
	}

	w.Write(respBody)
}
