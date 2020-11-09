package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	if err := run(getRunAddr()); err != nil {
		log.Fatal(err)
	}
}

func run(addr string) error {
	srv := &http.Server{
		Addr: addr,
	}

	log.Println("Serving on", srv.Addr)
	return srv.ListenAndServe()
}

func getRunAddr() string {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	return net.JoinHostPort("", port)
}
