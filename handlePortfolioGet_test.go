package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
	log "github.com/sirupsen/logrus"
)

func TestGetUserPortfolio(t *testing.T) {
	is := is.New(t)
	s := NewServer(ServerConfig{
		db:     NewTestDB(),
		logger: log.New(),
	})

	testCases := []struct {
		userid string
		code   int
	}{
		{"user1369", http.StatusOK},
		{"foobar", http.StatusBadRequest},
	}

	handler := s.routes()
	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", tc.userid), nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)
		is.Equal(tc.code, rr.Code)
	}
}
