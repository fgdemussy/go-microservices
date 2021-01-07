package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fgdemussy/go-microservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-ms", log.LstdFlags)

	// create handlers
	ph := handlers.NewProducts(l)

	// create a new serveMux and register the handlers
	port := "9000"
	addr := fmt.Sprintf(":%s", port)
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	srv := &http.Server{
		Handler:  sm,
		Addr:     addr,
		ErrorLog: l,
	}

	go func() {
		l.Printf("Starting server on %s", addr)
		err := srv.ListenAndServe()
		if err != nil {
			l.Printf("Server is shutting down: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully stop server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until signal received
	sig := <-c
	l.Println("got signal: ", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
