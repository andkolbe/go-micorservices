package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HandleFunc is a convenience method on the Go http package
// it registers a function to a path on the DefaultServeMux
// the DefaultServeMux is a http handler. Everything related to the server in Go is a http handler

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	// read everything that was in the request body into the variable d
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// http has an Error method built into it
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}
		
	fmt.Fprintf(rw, "Hello %s", d)
}

// ServeMux is responsible for redirecting paths. You map a function at a path and ServeMux will determine which function gets executed
// in the http package there is a DefaultServeMux. If you don't set anything else up, this is what is used

// a Handler in Go is just an interface
// any struct the has the method of ServeHTTP(ResponseWriter, *Request) implements the interface Handler