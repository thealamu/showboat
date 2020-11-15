package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handlePortfolioGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["id"]
		s.logger.Infof("Getting portfolio for user '%s'", userID)

		portfolio, err := s.db.GetPortfolio(r.Context(), userID)
		if err != nil {
			s.Error(w, err, http.StatusInternalServerError)
			return
		}
		s.JSON(w, portfolio, http.StatusOK)
	}
}
