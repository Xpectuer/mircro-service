package main

import (
	"context"
	"log"
	"my-simple-server/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)
	/**
	Serve Mux  is a map spcifies
	the routers and handler funcs
	*/
	sm := mux.NewRouter()
	log.Println("Registering Handlers... ")
	//sm.Handle("/products", ph)
	// In go public function started with Capital letter

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProducts)
	postRouter.Use(ph.MiddlewareProductValidation)
	// sm.Handle("/", ph)
	// sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:        ":9091",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
		// timeout for request reader
		ReadTimeout: 1 * time.Second,
		// timeout for response writer
		WriteTimeout: 1 * time.Second,
	}
	// graceful shutdown

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// specifies the channel
	sigChan := make(chan os.Signal)

	// relay the incoming signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// consume
	sig := <-sigChan
	l.Println("Received terminate, gracefully shutdown", sig)

	// allow handlers to gracefully shutdown in 30s
	// after 30s , forcefully close it
	// In golang 1.15 Version, cancel func cannot be swallowed
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)

}
