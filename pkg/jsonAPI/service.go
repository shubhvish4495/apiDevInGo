package jsonAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/shubhvish4495/apiDevInGo/pkg/service"
)

//getUserData gets data per each user
//userID - Id of the user we want to fetch data
func getUserData(userID int) (*User, error) {

	var srvrResp User

	resp, err := service.DoGetRequest(fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", userID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:

		defer resp.Body.Close()
		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct ", err)
			return nil, errors.New("can't unmarshal response into struct")
		}

	case http.StatusNotFound:
		return nil, errors.New("requested resource not found")

	default:
		return nil, errors.New("internal server error")
	}

	return &srvrResp, nil
}

func listPosts() ([]Post, error) {

	srvrResp := []Post{}

	resp, err := service.DoGetRequest("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:

		defer resp.Body.Close()
		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct ", err)
			return nil, errors.New("can't unmarshal response into struct")
		}

	case http.StatusNotFound:
		return nil, errors.New("requested resource not found")

	default:
		return nil, errors.New("internal server error")
	}

	return srvrResp, nil
}

func getComment(commentID int) (*Comment, error) {

	srvrResp := Comment{}

	resp, err := service.DoGetRequest(fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/%d", commentID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:

		defer resp.Body.Close()
		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct ", err)
			return nil, errors.New("can't unmarshal response into struct")
		}

	case http.StatusNotFound:
		return nil, errors.New("requested resource not found")

	default:
		return nil, errors.New("internal server error")
	}

	return &srvrResp, nil
}

func listComments() ([]Comment, error) {

	srvrResp := []Comment{}

	resp, err := service.DoGetRequest("https://jsonplaceholder.typicode.com/comments")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:

		defer resp.Body.Close()
		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct ", err)
			return nil, errors.New("can't unmarshal response into struct")
		}

	case http.StatusNotFound:
		return nil, errors.New("requested resource not found")

	default:
		return nil, errors.New("internal server error")
	}

	return srvrResp, nil
}

func showPost(postID int) (*Post, error) {

	srvrResp := Post{}

	resp, err := service.DoGetRequest(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:

		defer resp.Body.Close()
		err := json.NewDecoder(resp.Body).Decode(&srvrResp)

		if err != nil {
			log.Println("Error while decoding struct ", err)
			return nil, errors.New("can't unmarshal response into struct")
		}

	case http.StatusNotFound:
		return nil, errors.New("requested resource not found")

	default:
		return nil, errors.New("internal server error")
	}

	return &srvrResp, nil
}
