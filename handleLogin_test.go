package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/matryer/is"
	log "github.com/sirupsen/logrus"
)

func TestHandleLogin(t *testing.T) {
	is := is.New(t)

	type loginRequest struct {
		UserID   string `json:"userid"`
		Password string `json:"password"`
	}

	type loginResponse struct {
		Token string `json:"token"`
	}

	srv := NewServer(ServerConfig{
		db:     NewTestDB(),
		logger: log.New(),
	})

	testCases := []struct {
		desc     string
		userid   string
		password string
		code     int
	}{
		{"GoodParams", "user1369", "user1369Password", http.StatusOK},
		{"WrongPassword", "user1369", "foobarpassword", http.StatusUnauthorized},
		{"EmptyPassword", "user", "", http.StatusBadRequest},
		{"EmptyUserID", "", "pwd", http.StatusBadRequest},
		{"UserDoesNotExist", "notexist", "foob", http.StatusUnauthorized},
		{"BadUserID", "foo-bar", "foob", http.StatusBadRequest},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			body := strings.NewReader(fmt.Sprintf(`{"userid": "%s", "password": "%s"}`, tc.userid, tc.password))
			req := httptest.NewRequest(http.MethodPost, "/login", body)
			rr := httptest.NewRecorder()

			srv.routes().ServeHTTP(rr, req)

			is.Equal(tc.code, rr.Code)
		})
	}
}
