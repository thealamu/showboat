package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server is a http server that holds handler dependencies
type Server struct {
	*http.Server
	db         DB
	logger     *log.Logger
	hmacSecret string
}

func NewServer(cfg ServerConfig) *Server {
	s := &Server{
		&http.Server{
			Addr: cfg.Addr,
		},
		cfg.db,
		cfg.logger,
		cfg.hmacSecret,
	}
	s.Handler = s.routes()
	return s
}

func (s *Server) Start() error {
	//pre-flight checks
	if strings.TrimSpace(s.hmacSecret) == "" {
		panic("cannot start server without hmac secret")
	}

	if s.db == nil {
		panic("cannot start server without db")
	}

	if s.logger == nil {
		panic("cannot start server without a logger")
	}

	s.logger.Infof("Serving on %s", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf("/{id:%s}", UserIDFormat), s.handlePortfolioGet()).Methods(http.MethodGet)
	r.HandleFunc("/signup", s.handleSignup()).Methods(http.MethodPost)

	return r
}
