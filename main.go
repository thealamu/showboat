package main

import (
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	secret := os.Getenv("HMACSECRET")
	frontendURL := os.Getenv("FRONTEND")

	cfg := ServerConfig{
		logger:      logger,
		db:          &TestDB{},
		Addr:        getRunAddr(),
		hmacSecret:  secret,
		frontendURL: frontendURL,
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
