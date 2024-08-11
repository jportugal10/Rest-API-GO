package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndtoEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostHandler() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuite) TestPostNoResult() {
	c := http.Client{}
	r, _ := c.Get("http://localhost:8080/post/5336")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := io.ReadAll(r.Body)
	s.JSONEq(`{"status":"ok", "data":[]}`, string(b))
}
