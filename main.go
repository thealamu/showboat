package main

import (
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

var debug = os.Getenv("DEBUG") == "1"

func main() {
	logger := log.New()
	if debug {
		logger.SetLevel(log.DebugLevel)
	}

	secret := os.Getenv("HMACSECRET")
	frontendURL := os.Getenv("FRONTEND")

	cfg := ServerConfig{
		logger:      logger,
		db:          NewTestDB(),
		Addr:        getRunAddr(),
		hmacSecret:  secret,
		frontendURL: frontendURL,
	}

	if debug {
		if err := cfg.Dump(os.Stdout); err != nil {
			log.Warn("Could not dump config")
		}
	}

	srv := NewServer(cfg)
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
