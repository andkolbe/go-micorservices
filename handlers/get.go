package handlers

import (
	"net/http"
	"github.com/andkolbe/go-microservices/data"
)

// swagger:route GET /products products listProducts

// getProducts returns the products from the data store
// reading JSON from our server
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()
	// convert the list of products to JSON to send back to the user
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
