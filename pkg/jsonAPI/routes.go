package jsonAPI

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	router.
		Path("/posts/{postID}").
		Methods(http.MethodGet).
		HandlerFunc(ShowPostHandler)

	router.
		Path("/users/{userID}").
		Methods(http.MethodGet).
		HandlerFunc(UserDataHandler)

}
