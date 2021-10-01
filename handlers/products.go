package handlers

import (
	"log"
	"net/http"

	"github.com/andkolbe/go-microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// respond back to the user with the list of products in the data package
func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	// convert the list of products to JSON to send back to the user
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// an encoder does the same thing as marshal but instead of returning a slice of data or an error, it's writing the output direct to an io writer
// the reason we want to use an encoder and write direct is because then we aren't having to buffer anything into memory
// we don't have to allocate memory for the data object. If you have a large json document, then that could be a real consideration 
// the encoder is also faster. Makes a big deal with microservices