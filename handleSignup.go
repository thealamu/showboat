package main

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleSignup() http.HandlerFunc {
	type signupRequest struct {
		UserID   string `json:"userid"`
		Password string `json:"password"`
	}

	type signupResponse struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var signupReq signupRequest
		err := json.NewDecoder(r.Body).Decode(&signupReq)
		if err != nil {
			s.Error(w, err, http.StatusBadRequest)
			return
		}

		s.logger.Infof("Signing up user '%s'", signupReq.UserID)
		token, err := s.signupUser(r.Context(), signupReq.UserID, signupReq.Password)
		if err != nil {
			s.Error(w, err, http.StatusBadRequest)
			return
		}

		s.logger.Info("Successfully signed user up")
		s.JSON(w, signupResponse{token}, http.StatusCreated)
	}

}
