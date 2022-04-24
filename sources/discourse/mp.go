package discourse

import (
	"encoding/json"
	"enterpret/dataaccess"
	"enterpret/models"
	sourceInterface "enterpret/sources/interface"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type discourseMessageProcessor struct {
	store  dataaccess.DataStore
	logger *log.Logger
	client http.Client
}

func NewDiscourseMessageProcessor(store dataaccess.DataStore) sourceInterface.MessageProcessor {
	return &discourseMessageProcessor{store: store,
		logger: log.New(os.Stdout, "discourseMessageProcessor: ", 1),
		client: http.Client{
			Transport: http.DefaultTransport,
			Timeout:   time.Minute,
		}}
}

func (p *discourseMessageProcessor) fetchMessages(params models.Params) (models.DiscourseSearchResponse, error) {

	parse, err := url.Parse("https://meta.discourse.org/search.json?q=after%3A2021-01-01+before%3A2021-02-20")
	if err != nil {
		p.logger.Println("Error: ", err.Error())
		return models.DiscourseSearchResponse{}, err
	}

	//m := map[string][]string{}
	//
	//if params.SearchQuery != "" {
	//	m["q"] = []string{params.SearchQuery}
	//}
	//
	//if params.Before != nil {
	//	year, month, day := params.Before.Date()
	//	m["before"] = []string{fmt.Sprintf("%d-%d-%d", year, month, day)}
	//}
	//
	//if params.Since != nil {
	//	year, month, day := params.Since.Date()
	//	m["after"] = []string{fmt.Sprintf("%d-%d-%d", year, month, day)}
	//}

	do, err := p.client.Do(&http.Request{
		Method:           http.MethodGet,
		URL:              parse,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             "",
		Form:             nil,
		//Form:             m,
		PostForm:      nil,
		MultipartForm: nil,
		Trailer:       nil,
		RemoteAddr:    "",
		RequestURI:    "",
		TLS:           nil,
		Cancel:        nil,
		Response:      nil,
	})
	if err != nil {
		return models.DiscourseSearchResponse{}, err
	}

	fmt.Println(do.Body)

	var searchResp = models.DiscourseSearchResponse{}
	err = json.NewDecoder(do.Body).Decode(&searchResp)
	if err != nil {
		p.logger.Println(err.Error())
		return models.DiscourseSearchResponse{}, err
	}

	return searchResp, nil
}

func (p *discourseMessageProcessor) fetchIndividualMessage(id string) (models.DiscoursePost, error) {
	return models.DiscoursePost{}, nil
}

func (p *discourseMessageProcessor) FetchAndStoreMessages(params models.Params, tenant string) ([]models.MessageIngest, error) {

	searchResp, err := p.fetchMessages(params)
	if err != nil {
		return nil, err
	}

	for _, post := range searchResp.Posts {
		mI := models.MessageIngest{
			Meta: models.Meta{
				Tenant:       tenant,
				Source:       "discourse",
				CreationTime: post.CreatedAt,
				User:         post.Username,
				ID:           fmt.Sprint(post.Id),
				Blob:         nil,
			},
			Data: models.Data{
				Message: post.Blurb,
			},
		}

		err := p.store.Store(mI)
		if err != nil {
			p.logger.Println("Error: ", err.Error())
		}
	}
	return []models.MessageIngest{}, nil
}

func (p *discourseMessageProcessor) IngestMessage(blob []byte, tenant string) (models.MessageIngest, error) {
	var searchResp = models.DiscourseSearchResponse{}
	err := json.Unmarshal(blob, &searchResp)
	if err != nil {
		p.logger.Println(err.Error())
		return models.MessageIngest{}, err
	}

	for _, post := range searchResp.Posts {
		mI := models.MessageIngest{
			Meta: models.Meta{
				Tenant:       tenant,
				Source:       "discourse",
				CreationTime: post.CreatedAt,
				User:         post.Username,
				ID:           fmt.Sprint(post.Id),
				Blob:         nil,
			},
			Data: models.Data{
				Message: post.Blurb,
			},
		}

		err := p.store.Store(mI)
		if err != nil {
			p.logger.Println("Error: ", err.Error())
		}
	}

	return models.MessageIngest{}, nil
}

func (p *discourseMessageProcessor) IsSimilar(blob []byte, ingest models.MessageIngest) (bool, error) {
	return false, nil
}
