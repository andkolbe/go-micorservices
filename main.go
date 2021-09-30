package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andkolbe/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// ListenAndServe contructs a http server and registers a default handler to it. If a handler is not specified, it uses DefaultServeMux
	http.ListenAndServe(":9090", sm)
}

