package models

import "time"

//MessageIngest stores the ingested message from various sources
type MessageIngest struct {
	Meta Meta
	Data Data
}

//Meta stores the metadata part of all message ingests
type Meta struct {
	Tenant       string
	Source       string
	CreationTime time.Time
	User         string
	ID           string
	Blob         []byte
}

//Data stores the data part of all message ingests
type Data struct {
	Message string
}

//Params are used by the pull workflow to fetch messages from a source
type Params struct {
	Since, Before *time.Time
	SearchQuery   string
	Source        string
}
