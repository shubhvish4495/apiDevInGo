package jsonAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func showData(postID int) (error, ShowResponse) {

	srvrResp := ShowResponse{}

	req, err :=
		http.NewRequest(http.MethodGet, fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID), nil)
	if err != nil {
		return err, srvrResp
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, srvrResp
	}

	switch resp.StatusCode {
	case http.StatusOK:

		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct %v", err)
			return errors.New("can't unmarshal response into struct"), srvrResp
		}

	case http.StatusNotFound:
		return errors.New("requested resource not found"), srvrResp

	default:
		return errors.New("internal server error"), srvrResp
	}

	return nil, srvrResp
}
