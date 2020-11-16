package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
	log "github.com/sirupsen/logrus"
)

func TestHandleUidForToken(t *testing.T) {
	is := is.New(t)
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	srv := NewServer(ServerConfig{
		logger:     logger,
		db:         &TestDB{},
		hmacSecret: "testSecret",
	})

	goodTestToken, err := srv.createToken("testuser")
	is.NoErr(err)
	badTestToken := strings.ReplaceAll(goodTestToken, "e", "a")

	testCases := []struct {
		desc   string
		token  string
		code   int
		output string
	}{
		{"GoodTestToken", goodTestToken, http.StatusOK, `{"userid":"testuser"}`},
		{"BadTestToken", badTestToken, http.StatusUnauthorized, ""},
		{"EmptyTestToken", "", http.StatusUnauthorized, ""},
	}

	handler := srv.routes()
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/i/user", nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.token))
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			is.Equal(tc.code, rr.Code)
			if tc.output != "" {
				respBody, err := ioutil.ReadAll(rr.Body)
				is.NoErr(err)
				respStr := strings.TrimSpace(string(respBody))
				is.Equal(tc.output, respStr)
			}
		})
	}
}
