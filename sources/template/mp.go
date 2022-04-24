package template

import (
	"enterpret/models"
	sourceInterface "enterpret/sources/interface"
)

type template struct{}

func (t *template) FetchAndStoreMessages(params models.Params, tenant string) ([]models.MessageIngest, error) {
	return []models.MessageIngest{}, nil
}

func (t *template) IngestMessage(blob []byte, tenant string) (models.MessageIngest, error) {
	return models.MessageIngest{}, nil
}

func (t *template) IsSimilar(blob []byte, ingest models.MessageIngest) (bool, error) {
	return false, nil
}

func NewTemplateSourceProcessor() sourceInterface.MessageProcessor {
	return &template{}
}
