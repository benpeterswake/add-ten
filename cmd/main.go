package main

import (
	"log"
	"net/http"
	"time"

	function "github.com/benpeterswake/add-ten"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{num}", function.RootHandler).Methods("GET")

	srv := &http.Server{
		Handler: handlers.CORS()(r),
		Addr:    "127.0.0.1:8008",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Running server on port 8008")
	log.Fatal(srv.ListenAndServe())
}
