package sourceInterface

import "enterpret/models"

type MessageProcessor interface {
	FetchAndStoreMessages(params models.Params, tenant string) ([]models.MessageIngest, error)
	IngestMessage(blob []byte, tenant string) (models.MessageIngest, error)
	IsSimilar(blob []byte, ingest models.MessageIngest) (bool, error)
}

type FetchMessage interface {
}
