package jsonAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/shubhvish4495/apiDevInGo/pkg/service"
)

func showData(postID int) (*ShowResponse, error) {

	srvrResp := ShowResponse{}

	resp, err := service.DoGetRequest(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:

		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct %v", err)
			return nil, errors.New("can't unmarshal response into struct")
		}

	case http.StatusNotFound:
		return nil, errors.New("requested resource not found")

	default:
		return nil, errors.New("internal server error")
	}

	return &srvrResp, nil
}
