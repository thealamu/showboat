package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Server is a http server that holds handler dependencies
type Server struct {
	*http.Server
	db DB
}

func (s *Server) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf("/{id:%s}", UserIDFormat), s.handlePortfolioGet()).Methods(http.MethodGet)

	return r
}
