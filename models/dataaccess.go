package models

import "time"

//FeedbackIngest stores the ingested message from various sources
type FeedbackIngest struct {
	Meta Meta
	Data Data
}

//Meta stores the metadata part of all message ingests
type Meta struct {
	Tenant       string
	Source       string
	Language     string
	Params       Params
	CreationTime time.Time
	User         string
	ID           string
	Attributes   interface{}
}

//Data stores the data part of all message ingests
type Data struct {
	Message interface{}
}

//Params are used by the pull workflow to fetch messages from a source
type Params struct {
	Since, Before            *time.Time
	SearchQuery              string
	SourceSpecificParameters map[string]interface{}
}
