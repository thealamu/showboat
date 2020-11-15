package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
	log "github.com/sirupsen/logrus"
)

func TestGetUserPortfolio(t *testing.T) {
	is := is.New(t)
	s := &Server{
		db:     TestDB{},
		logger: log.New(),
	}

	handler := s.routes()
	req := httptest.NewRequest(http.MethodGet, "/user1369", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	rrBody, err := ioutil.ReadAll(rr.Body)
	is.NoErr(err)

	expected, err := json.Marshal(getTestPortfolio())
	is.NoErr(err)

	rrBodyStr := strings.TrimSpace(string(rrBody))
	expectedStr := strings.TrimSpace(string(expected))

	is.Equal(expectedStr, rrBodyStr)
}
