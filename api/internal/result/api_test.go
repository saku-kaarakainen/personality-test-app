package result

import (
	"net/http"
	"testing"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/test"
)

func TestAPI(t *testing.T) {
	router := test.MockGinRouter()
	// header := http.Header{}

	// define in service_test.go
	repo := mockRepository{
		// items: []entity.Question{
		// 	{
		// 		Id:          "1",
		// 		Text:        "text",
		// 		Description: "desc",
		// 		Answers: []entity.Answer{{
		// 			Id:    "1_1",
		// 			Score: [2]int{0, 0},
		// 			Label: "label",
		// 		}},
		// 	},
		// },
	}

	RegisterHandlers(router, NewService(repo, test.MockLoader{
		RetBytes: []byte{},
	}))

	tests := []test.APITestCase{
		{
			Name:         "result: calculate result",
			Method:       "GET",
			URL:          "/result/calculate?q[0]=1&q[1]=2&q[2]=3",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
		{
			Name:         "result: calculate, no params",
			Method:       "GET",
			URL:          "/result/calculate",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: "",
		},
		{
			Name:         "result: unknown route",
			Method:       "GET",
			URL:          "/result",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusNotFound,
			WantResponse: "",
		},
		{
			Name:         "result: wrong verb",
			Method:       "POST",
			URL:          "/result/calculate",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusNotFound,
			WantResponse: "",
		},
	}

	for _, testCase := range tests {
		test.Endpoint(t, router, testCase)
	}
}
