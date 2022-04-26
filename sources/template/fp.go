package template

import (
	"enterpret/models"
	sourceInterface "enterpret/sources/sourceinterface"
)

type template struct{}

func (t template) FetchAndStoreFeedbacks(params models.Params, tenant string) ([]models.FeedbackIngest, error) {
	//TODO implement me
	panic("implement me")
}

func (t template) IngestAndStoreFeedback(blob []byte, tenant string) (models.FeedbackIngest, error) {
	//TODO implement me
	panic("implement me")
}

func NewTemplateSourceProcessor() sourceInterface.FeedbackProcessor {
	return &template{}
}
