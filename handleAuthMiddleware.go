package main

import (
	"errors"
	"net/http"
	"strings"
)

func (s *Server) handleAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := s.tokenFromHeader(r)
		s.logger.Debugf("Stripped token %s", token)
		if token == "" {
			s.Error(w, errors.New("got no token"), http.StatusUnauthorized)
			return
		}
		userid, err := s.claimFromToken(token)
		if err != nil {
			statusCode := http.StatusUnauthorized
			if err == ErrBadClaimType {
				statusCode = http.StatusInternalServerError
			}
			s.logger.Warn(err)
			s.Error(w, err, statusCode)
			return
		}

		newCtx := s.putUidInContext(r.Context(), userid)
		next.ServeHTTP(w, r.WithContext(newCtx))
	}
}

func (s *Server) tokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	s.logger.Debug(authHeader)
	return strings.TrimPrefix(authHeader, "Bearer ")
}
