package main

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (s *Server) handleLogin() http.HandlerFunc {
	type loginRequest struct {
		UserID   string `json:"userid"`
		Password string `json:"password"`
	}

	type loginResponse struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var loginReq loginRequest
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			s.Error(w, err, http.StatusBadRequest)
			return
		}

		s.logger.Infof("Logging in user '%s'", loginReq.UserID)
		token, err := s.loginUser(r.Context(), loginReq.UserID, loginReq.Password)
		if err != nil {
			statusCode := http.StatusUnauthorized
			if errors.Cause(err) == ErrBadCredInput {
				statusCode = http.StatusBadRequest
			}
			s.Error(w, err, statusCode)
			return
		}

		s.logger.Info("Successfully logged user in")
		s.JSON(w, loginResponse{token}, http.StatusOK)
	}
}
