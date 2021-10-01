package handlers

import (
	"log"
	"net/http"

	"github.com/andkolbe/go-microservices/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


// an encoder does the same thing as marshal but instead of returning a slice of data or an error, it's writing the output direct to an io writer
// the reason we want to use an encoder and write direct is because then we aren't having to buffer anything into memory
// we don't have to allocate memory for the data object. If you have a large json document, then that could be a real consideration 
// the encoder is also faster. Makes a big deal with microservices

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
// writing JSON back to the user
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// if the incoming request is a GET, display the products on the screen
	// respond back to the user with the list of products in the data package
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}


	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
// reading JSON from our server
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()
	// convert the list of products to JSON to send back to the user
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	// create new product object
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		// we got a bad request
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest) 
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}