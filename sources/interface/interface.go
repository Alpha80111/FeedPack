package sourceInterface

import "enterpret/models"

//go:generate mockgen -package=mock -destination=mock/interface.go -source=interface.go FeedbackProcessor
type FeedbackProcessor interface {
	FetchAndStoreFeedbacks(params models.Params, tenant string) ([]models.FeedbackIngest, error)
	IngestAndStoreFeedback(blob []byte, tenant string) (models.FeedbackIngest, error)
}
