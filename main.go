package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shubhvish4495/apiDevInGo/pkg/jsonAPI"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/posts/{postId}", jsonAPI.ShowHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:4444",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on 4444")
	log.Fatal(srv.ListenAndServe())
}
