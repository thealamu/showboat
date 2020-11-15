package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server is a http server that holds handler dependencies
type Server struct {
	*http.Server
	db     DB
	logger *log.Logger
}

func (s *Server) Start() error {
	s.logger.Infof("Serving on %s", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf("/{id:%s}", UserIDFormat), s.handlePortfolioGet()).Methods(http.MethodGet)
	r.HandleFunc("/signup", s.handleSignup()).Methods(http.MethodPost)

	return r
}
