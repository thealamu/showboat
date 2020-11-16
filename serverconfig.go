package main

import (
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"
)

//ServerConfig defines the dependencies of the http server
type ServerConfig struct {
	logger      *log.Logger
	db          DB
	Addr        string
	hmacSecret  string
	frontendURL string
}

func (s ServerConfig) Dump(w io.Writer) error {
	_, err := fmt.Fprintf(w, fmt.Sprintf("%+v\n", s))
	return err
}
