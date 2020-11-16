package main

import "net/http"

func (s *Server) handleUidForToken() http.HandlerFunc {
	type response struct {
		UserID string `json:"userid"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		userid, err := s.uidFromContext(r.Context())
		if err != nil {
			s.Error(w, err, http.StatusInternalServerError)
			return
		}
		s.JSON(w, response{userid}, http.StatusOK)
	}
}
