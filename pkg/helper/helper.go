package helper

import "net/http"

//ReturnSuccessResponse returns success response
func ReturnSuccessResponse(v interface{}, rw http.ResponseWriter) {}

//ReturnErrResponse returns error response
func ReturnErrResponse(e error, rw http.ResponseWriter) {}
