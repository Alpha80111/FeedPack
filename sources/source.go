package sources

import (
	"enterpret/dataaccess"
	"enterpret/sources/discourse"
	_interface "enterpret/sources/interface"
	"errors"
)

type SourceProcessor interface {
	GetProcessor(source string) (_interface.MessageProcessor, error)
}

type sourceProcessor struct {
	discourse _interface.MessageProcessor
}

func (s *sourceProcessor) GetProcessor(source string) (_interface.MessageProcessor, error) {
	switch source {
	case "discourse":
		return s.discourse, nil
	}

	return nil, errors.New("source message processor not found")
}

func NewSourceProcessor(store dataaccess.DataStore) SourceProcessor {
	dMP := discourse.NewDiscourseMessageProcessor(store)

	return &sourceProcessor{dMP}
}
