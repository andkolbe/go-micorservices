package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// HandleFunc is a convenience method on the Go http package
	// it registers a function to a path on the DefaultServeMux
	// the DefaultServeMux is a http handler. Everything related to the server in Go is a http handler
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		// read everything that was in the request body into the variable d
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
		
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	// ListenAndServe contructs a http server and registers a default handler to it. If a handler is not specified, it uses DefaultServeMux
	http.ListenAndServe(":9090", nil)
}

// ServeMux is responsible for redirecting paths. You map a function at a path and ServeMux will determine which function gets executed
// in the http package there is a DefaultServeMux. If you don't set anything else up, this is what is used

// a Handler in Go is just an interface
// any struct the has the method of ServeHTTP(ResponseWriter, *Request) implements the interface Handler