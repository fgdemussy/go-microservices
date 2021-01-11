package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/fgdemussy/go-microservices/data"
	"github.com/gorilla/mux"
)

// Products defines a handler for products related requests
type Products struct {
	l *log.Logger
}

// KeyProduct is a key to store Product in request context
type KeyProduct struct {}


// NewProducts returns a new Products handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts gets a list of products from data store and returns them in json format
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Could not encode the product list", http.StatusInternalServerError)
	}
}

// AddProduct adds a product to the list
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Handle GET Products")
	
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

// UpdateProduct updates a given product
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to parse ID, is that an integer?", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Product", id)
	
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("got prod from context %v", prod)

	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// MiddlewareProductValidation extracts a Product from the request body
func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler  {
	
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}