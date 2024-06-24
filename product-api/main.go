package main

import (
	"context"
	"gRPC-go/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server.")

func main() {

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/products/", ph)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Printf("Starting Server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block
	sig := <-c
	log.Printf("Got signal: %v, exiting", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
