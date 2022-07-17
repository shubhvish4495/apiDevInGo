package jsonAPI

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/posts/{postId}", ShowPostHandler)
	router.HandleFunc("/users/{userID}", UserDataHandler)
}
