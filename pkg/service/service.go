package service

import "net/http"

func doServiceRequest(url, method string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func DoGetRequest(url string) (*http.Response, error) {
	return doServiceRequest(url, http.MethodGet, nil)
}
