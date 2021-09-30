package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andkolbe/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create a http server
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,

	}

	// ListenAndServe contructs a http server and registers a default handler to it. If a handler is not specified, it uses DefaultServeMux
	s.ListenAndServe()
}

