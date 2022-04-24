package dataaccess

import (
	"enterpret/models"
	"errors"
	"fmt"
)

type DataStore interface {
	Store(ingest models.MessageIngest) error
	FetchMessages(opts ...Opts) ([]models.MessageIngest, error)
}

type params struct {
	source string
	tenant string
}

type Opts func(params *params)

func GetTenantOption(tenant string) Opts {
	return func(params *params) {
		params.tenant = tenant
	}
}

func GetSourceOption(source string) Opts {
	return func(params *params) {
		params.source = source
	}
}

type dataStore struct {
	store map[string]map[string]map[string]models.MessageIngest //map[tenant]map[source]map[post_id]models.MessageIngest
}

func (d *dataStore) Store(ingest models.MessageIngest) error {
	fmt.Printf("Received message: %s, %s, %s\n", ingest.Meta.Tenant, ingest.Meta.Source, ingest.Meta.ID)
	if _, ok := d.store[ingest.Meta.Tenant]; !ok {
		d.store[ingest.Meta.Tenant] = map[string]map[string]models.MessageIngest{}
	}
	if _, ok := d.store[ingest.Meta.Tenant][ingest.Meta.Source]; !ok {
		d.store[ingest.Meta.Tenant][ingest.Meta.Source] = map[string]models.MessageIngest{}
	}
	if _, ok := d.store[ingest.Meta.Tenant][ingest.Meta.Source][ingest.Meta.ID]; !ok {
		d.store[ingest.Meta.Tenant][ingest.Meta.Source][ingest.Meta.ID] = ingest
	}
	return nil
}

func (d *dataStore) FetchMessages(opts ...Opts) ([]models.MessageIngest, error) {
	p := params{
		source: "",
		tenant: "",
	}
	for _, opt := range opts {
		opt(&p)
	}
	if p.source == "" && p.tenant == "" {
		return []models.MessageIngest{}, errors.New("no valid options passed")
	}
	var messages []models.MessageIngest
	if p.tenant != "" {
		if p.source != "" {
			for i := range d.store[p.tenant][p.source] {
				messages = append(messages, d.store[p.tenant][p.source][i])
			}
		}
	}
	return messages, nil
}

func NewDataStore() DataStore {
	return &dataStore{store: map[string]map[string]map[string]models.MessageIngest{}}
}
