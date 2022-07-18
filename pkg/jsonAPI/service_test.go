package jsonAPI

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shubhvish4495/apiDevInGo/pkg/helper"
	"github.com/shubhvish4495/apiDevInGo/pkg/service"
)

func TestGetUserData(t *testing.T) {
	// jsonResp := `{"id":1,"name":"Test User","email":"testUser@test.com"}`

	testCases := []helper.TestCaseRestAPI{
		// {
		// 	Name: "Success Case",
		// 	ExpectedResp: &User{
		// 		ID:    1,
		// 		Name:  "Test User",
		// 		Email: "testUser@test.com",
		// 	},
		// 	MockDoFunc: func(req *http.Request) (*http.Response, error) {
		// 		return &http.Response{
		// 			StatusCode: http.StatusOK,
		// 			Body:       io.NopCloser(bytes.NewReader([]byte(jsonResp))),
		// 		}, nil
		// 	},
		// },
		{
			Name:         "UserId not found",
			ExpectedResp: errors.New("requested resource not found"),
			MockDoFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusNotFound,
				}, nil
			},
		},
		// {
		// 	Name:         "Internal Server Error",
		// 	ExpectedResp: errors.New("internal server error"),
		// 	MockDoFunc: func(req *http.Request) (*http.Response, error) {
		// 		return &http.Response{
		// 			StatusCode: http.StatusInternalServerError,
		// 		}, nil
		// 	},
		// },
		// {
		// 	Name:         "Bad Response Data",
		// 	ExpectedResp: errors.New("can't unmarshal response into struct"),
		// 	MockDoFunc: func(req *http.Request) (*http.Response, error) {
		// 		return &http.Response{
		// 			StatusCode: http.StatusOK,
		// 			Body:       io.NopCloser(bytes.NewReader([]byte("["))),
		// 		}, nil
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			service.Client = &helper.MockClient{DoFunc: tc.MockDoFunc}
			data, err := getUserData(1)
			if val, ok := tc.ExpectedResp.(*User); !ok {
				assert.EqualValues(t, tc.ExpectedResp.(error), err)
			} else {
				assert.EqualValues(t, val, data)
			}
		})
	}

}
