package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/shubhvish4495/apiDevInGo/pkg/jsonAPI"
	"github.com/shubhvish4495/apiDevInGo/pkg/middleware"
)

func main() {

	r := mux.NewRouter()

	//setup logger middleware
	r.Use(middleware.LogRequest)

	//register routes
	jsonAPI.RegisterRoutes(r)

	//initialize server
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
