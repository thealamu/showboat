package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	is "github.com/matryer/is"
	log "github.com/sirupsen/logrus"
)

func TestHandleSignup(t *testing.T) {
	is := is.New(t)

	type signupRequest struct {
		UserID   string `json:"userid"`
		Password string `json:"password"`
	}

	type signupResponse struct {
		Token string `json:"token"`
	}

	srv := NewServer(ServerConfig{
		db:     &TestDB{},
		logger: log.New(),
	})

	testCases := []struct {
		desc     string
		userid   string
		password string
		code     int
	}{
		{"GoodParams", "foobar", "foobarpassword", http.StatusCreated},
		{"EmptyPassword", "user", "", http.StatusBadRequest},
		{"EmptyUserID", "", "pwd", http.StatusBadRequest},
		{"UserAlreadyExists", "foobar", "foob", http.StatusBadRequest},
		{"BadUserID", "foo-bar", "foob", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			body := strings.NewReader(fmt.Sprintf(`{"userid": "%s", "password": "%s"}`, tc.userid, tc.password))
			req := httptest.NewRequest(http.MethodPost, "/signup", body)
			rr := httptest.NewRecorder()

			srv.routes().ServeHTTP(rr, req)

			is.Equal(tc.code, rr.Code)
		})
	}
}
