package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestGetUserPortfolio(t *testing.T) {
	is := is.New(t)
	s := &Server{
		db: TestDB{},
	}

	handler := s.routes()
	req := httptest.NewRequest(http.MethodGet, "/user1369", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	rrBody, err := ioutil.ReadAll(rr.Body)
	is.NoErr(err)

	expected, err := json.Marshal(getTestPortfolio())
	is.NoErr(err)

	is.Equal(string(expected), string(rrBody))
}
