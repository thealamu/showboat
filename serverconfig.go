package main

import log "github.com/sirupsen/logrus"

//ServerConfig defines the dependencies of the http server
type ServerConfig struct {
	logger      *log.Logger
	db          DB
	Addr        string
	hmacSecret  string
	frontendURL string
}
