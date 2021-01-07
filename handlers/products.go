package handlers

import (
	"encoding/json"
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
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Could not encode the product list", http.StatusInternalServerError)
		return
	}
	w.Write(d)
}
