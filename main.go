package main

import (
	"net"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	srv := &Server{
		&http.Server{
			Addr: getRunAddr(),
		},
		TestDB{},
		logger,
	}
	srv.Handler = srv.routes()

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

func getRunAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	return net.JoinHostPort("", port)
}
