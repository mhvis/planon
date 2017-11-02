package main

import (
	"context"
	//	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server http.Server

func NewServer() *Server {
	r := mux.NewRouter()

	r.Handle("/bookings", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hallo"))
	}))
	//r.Handle("/bookings", BookingsHandler(session)).Methods("GET")
	//r.Handle("/book", BookHandler(session)).Methods("GET")

	h := &http.Server{
		Addr:    ":3000",
		Handler: handlers.LoggingHandler(os.Stdout, r),
	}

	go func() {
		log.Println("Listening on http://0.0.0.0:3000")
		err := h.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return (*Server)(h)
}

func (s *Server) Shutdown() {
	log.Println("Shutting down the server...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	(*http.Server)(s).Shutdown(ctx)

	log.Println("Server gracefully stopped")
}

func BookingsHandler() {

}
