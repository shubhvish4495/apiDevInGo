package jsonAPI

import (
	"encoding/json"
	"net/http"
)

func ShowHandler(rw http.ResponseWriter, r *http.Request) {

	err, resp := showData(2)

	if err != nil {
		rw.Write([]byte(err.Error()))
	}

	json.NewEncoder(rw).Encode(resp)

	// rw.Write([]byte("Hello World!"))
}
