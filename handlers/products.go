package handlers

import (
	"log"
	"net/http"

	"github.com/fgdemussy/go-microservices/data"
)

// Products defines a handler for products related requests
type Products struct {
	l *log.Logger
}

// NewProducts returns a new Products handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Could not encode the product list", http.StatusInternalServerError)
	}
}
