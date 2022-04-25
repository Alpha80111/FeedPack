package sourceInterface

import "enterpret/models"

type FeedbackProcessor interface {
	FetchAndStoreFeedbacks(params models.Params, tenant string) ([]models.FeedbackIngest, error)
	IngestAndStoreFeedback(blob []byte, tenant string) (models.FeedbackIngest, error)
}
