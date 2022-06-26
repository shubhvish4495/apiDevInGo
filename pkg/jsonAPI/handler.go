package jsonAPI

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shubhvish4495/apiDevInGo/pkg/helper"
)

var defaultPostID = 2

//ShowHandler returns the showData
func ShowHandler(rw http.ResponseWriter, r *http.Request) {
	var postID int

	if postVar, err := strconv.Atoi(mux.Vars(r)["postId"]); err == nil {
		postID = postVar
	} else {
		postID = defaultPostID
	}

	resp, err := showData(postID)
	if err != nil {
		rw.Write([]byte(err.Error()))
	}

	helper.ReturnJSONResponse(rw, resp)
}
