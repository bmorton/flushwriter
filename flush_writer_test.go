package flushwriter

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FlushWriterTestSuite struct {
	suite.Suite
}

func (suite *FlushWriterTestSuite) TestWrites() {
	w := httptest.NewRecorder()
	subject := New(w)
	subject.Write([]byte("HTTP/1.1 200 OK"))

	assert.Equal(suite.T(), "HTTP/1.1 200 OK", w.Body.String())
}

func (suite *FlushWriterTestSuite) TestWriteFlushesAutomatically() {
	w := httptest.NewRecorder()
	subject := New(w)
	subject.Write([]byte("test"))

	assert.True(suite.T(), w.Flushed)
}

func TestFlushWriterTestSuite(t *testing.T) {
	suite.Run(t, new(FlushWriterTestSuite))
}
