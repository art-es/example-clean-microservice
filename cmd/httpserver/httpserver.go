package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/art-es/example-clean-microservice/internal/infrastructure/httphandler"
)

var addr string

func main() {
	flag.StringVar(&addr, "addr", ":8080", "address of HTTP server")
	flag.Parse()

	if err := run(); err != nil {
		log.Panicf("[PANIC] main: run: %v\n", err)
	}
}

func run() error {
	mux, err := httphandler.NewServeMux()
	if err != nil {
		return err
	}

	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("[INFO] Server started on http://127.0.0.1%s\n", srv.Addr)
	return srv.ListenAndServe()
}
