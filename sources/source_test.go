package sources

import (
	"enterpret/dataaccess/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetProcessor(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := NewSourceProcessor(mock.NewMockDataStore(ctrl))

	fp, err := s.GetProcessor("discourse")
	assert.Nil(t, err)
	assert.NotNil(t, fp)

	fp, err = s.GetProcessor("error-source")
	assert.Nil(t, fp)
	assert.NotNil(t, err)
}
