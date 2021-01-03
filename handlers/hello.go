package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is just another handler with a logger injected
type Hello struct {
	l *log.Logger
}

// NewHello returns a Hello handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// Handle takes care of the request
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("handle hello requests")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "hello there '%s'", b)
}
