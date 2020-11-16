package main

import (
	"errors"
	"net/http"
	"strings"
)

func (s *Server) handlAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := s.tokenFromHeader(r)
		s.logger.Debugf("Stripped token %s", token)
		if token == "" {
			s.Error(w, errors.New("got no token"), http.StatusUnauthorized)
			return
		}
		userid, err := s.claimFromToken(token)
		if err != nil {
			s.logger.Warn(err)
			return
		}

		newCtx := s.putUidInContext(r.Context(), userid)
		next.ServeHTTP(w, r.WithContext(newCtx))
	}
}

func (s *Server) tokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authentication")
	return strings.TrimLeft("Bearer ", authHeader)
}
