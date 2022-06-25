package helper

import (
	"encoding/json"
	"net/http"
)

//ReturnJSONResponse returns JSON response
func ReturnJSONResponse(rw http.ResponseWriter, data interface{}) (int, error) {

	if data == nil {
		return 0, nil
	}

	rw.Header().Add("Content-Type", "application/json")
	strData, _ := json.Marshal(data)

	// rw.WriteHeader(http.StatusOK)
	return rw.Write(strData)

}

//ReturnErrResponse returns error response
func ReturnErrResponse(e error, rw http.ResponseWriter) {}
