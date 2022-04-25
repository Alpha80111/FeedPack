package template

import (
	"enterpret/models"
	sourceInterface "enterpret/sources/interface"
)

type template struct{}

func (t *template) FetchAndStoreFeedbacks(params models.Params, tenant string) ([]models.FeedbackIngest, error) {
	return []models.FeedbackIngest{}, nil
}

func (t *template) IngestAndStoreFeedback(blob []byte, tenant string) (models.FeedbackIngest, error) {
	return models.FeedbackIngest{}, nil
}

func (t *template) IsSimilar(blob []byte, ingest models.FeedbackIngest) (bool, error) {
	return false, nil
}

func NewTemplateSourceProcessor() sourceInterface.FeedbackProcessor {
	return &template{}
}
