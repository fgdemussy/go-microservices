package handlers

import (
	"io"
	"log"
	"net/http"
)

// Bazooka handles requests
type Bazooka struct {
	l *log.Logger
}

// NewBazooka returns a Bazooka handler
func NewBazooka(l *log.Logger) *Bazooka {
	return &Bazooka{l}
}

// Handle responds to http
func (b *Bazooka) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b.l.Printf("this is my bazooka, %v", b)
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, "Hey bazooka")
	return
}
