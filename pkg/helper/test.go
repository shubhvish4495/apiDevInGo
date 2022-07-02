package helper

import "net/http"

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

type MockResp struct {
	*http.Response
	error
}

type TestCaseRestAPI struct {
	Name         string
	ExpectedResp interface{}
	MockDoFunc   func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
