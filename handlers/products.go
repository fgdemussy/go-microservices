// Package handlers Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fgdemussy/go-microservices/data"
	"github.com/gorilla/mux"
)

// A list of products returned in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All Products found in the system
	// in: body
	Body []data.Product
}

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

// GetProducts returns a list of products
// swagger:route GET /products products listProducts
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//     Responses:
//     - 200: productsResponse
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
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
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}