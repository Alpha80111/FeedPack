package dataaccess

import (
	"enterpret/models"
	"errors"
	"fmt"
	"log"
	"os"
)

//DataStore interface that supports storing and fetching feedbacks
//go:generate mockgen -package=mock -destination=mock/dataaccess.go -source=dataaccess.go DataStore
type DataStore interface {
	Store(ingest models.FeedbackIngest) error
	FetchFeedbacks(tenant, source string, page, size int) ([]models.FeedbackIngest, error)
}

type params struct {
	source string
	tenant string
	page   int
	size   int
}

type dataStore struct {
	logger *log.Logger
	order  map[string]map[string][]string                         //map[tenant]map[source][]string(postIDs)
	store  map[string]map[string]map[string]models.FeedbackIngest //map[tenant]map[source]map[post_id]models.FeedbackIngest
}

//Store stores a message ingest passed to it
func (d *dataStore) Store(ingest models.FeedbackIngest) error {
	if ingest.Meta.Tenant == "" || ingest.Meta.Source == "" || ingest.Meta.ID == "" {
		return errors.New("invalid feedback, tenant, source and ID cannot be nil")
	}
	if _, ok := d.store[ingest.Meta.Tenant]; !ok {
		d.store[ingest.Meta.Tenant] = map[string]map[string]models.FeedbackIngest{}
		d.order[ingest.Meta.Tenant] = map[string][]string{}
	}
	if _, ok := d.store[ingest.Meta.Tenant][ingest.Meta.Source]; !ok {
		d.store[ingest.Meta.Tenant][ingest.Meta.Source] = map[string]models.FeedbackIngest{}
		d.order[ingest.Meta.Tenant][ingest.Meta.Source] = []string{}
	}
	if _, ok := d.store[ingest.Meta.Tenant][ingest.Meta.Source][ingest.Meta.ID]; !ok {
		d.store[ingest.Meta.Tenant][ingest.Meta.Source][ingest.Meta.ID] = ingest
		d.order[ingest.Meta.Tenant][ingest.Meta.Source] = append(d.order[ingest.Meta.Tenant][ingest.Meta.Source], ingest.Meta.ID)
		fmt.Printf("Storing message: %s, %s, %s\n", ingest.Meta.Tenant, ingest.Meta.Source, ingest.Meta.ID)
	} else {
		fmt.Printf("Duplicate message: %s, %s, %s\n", ingest.Meta.Tenant, ingest.Meta.Source, ingest.Meta.ID)
	}
	return nil
}

//FetchFeedbacks fetches messages within the parameters passed
func (d *dataStore) FetchFeedbacks(tenant, source string, page, size int) ([]models.FeedbackIngest, error) {

	if source == "" || tenant == "" || page <= 0 || size <= 0 {
		return []models.FeedbackIngest{}, errors.New("no valid options passed")
	}
	var messages []models.FeedbackIngest
	if tenant != "" {
		if source != "" {
			for i := (page - 1) * size; i < page*size && i < len(d.order[tenant][source]); i++ {
				messages = append(messages, d.store[tenant][source][d.order[tenant][source][i]])
			}
		}
	}
	return messages, nil
}

//NewDataStore initializes and returns a new DataStore
func NewDataStore() DataStore {
	return &dataStore{
		logger: log.New(os.Stdout, "Data Store: ", 1),
		store:  map[string]map[string]map[string]models.FeedbackIngest{},
		order:  map[string]map[string][]string{},
	}
}
