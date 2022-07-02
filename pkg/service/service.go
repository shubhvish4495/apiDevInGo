package service

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

func doServiceRequest(url, method string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DoGetRequest(url string) (*http.Response, error) {
	return doServiceRequest(url, http.MethodGet, nil)
}
