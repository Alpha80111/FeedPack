package sources

import (
	"enterpret/dataaccess"
	"enterpret/sources/discourse"
	_interface "enterpret/sources/interface"
	"errors"
)

const (
	DiscourseSourceKey = "discourse"
)

//go:generate mockgen -package=mock -destination=mock/source.go -source=source.go SourceProcessor
type SourceProcessor interface {
	GetProcessor(source string) (_interface.FeedbackProcessor, error)
}

type sourceProcessor struct {
	discourse _interface.FeedbackProcessor
}

func (s *sourceProcessor) GetProcessor(source string) (_interface.FeedbackProcessor, error) {
	switch source {
	case DiscourseSourceKey:
		return s.discourse, nil
	}

	return nil, errors.New("source message processor not found")
}

func NewSourceProcessor(store dataaccess.DataStore) SourceProcessor {
	dMP := discourse.NewDiscourseFeedbackProcessor(store)

	return &sourceProcessor{dMP}
}
