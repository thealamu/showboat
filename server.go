package main

import "net/http"

// Server is a http server that holds handler dependencies
type Server struct {
	*http.Server
	db DB
}

func (s *Server) routes() http.Handler {
	return nil
}
