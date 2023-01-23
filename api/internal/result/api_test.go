package result

import (
	"net/http"
	"testing"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/test"
	"github.com/stretchr/testify/mock"
)

func TestAPI(t *testing.T) {
	router := test.MockGinRouter()
	srvs := new(mockService)

	res := Result{entity.Result{
		Id:                    "1",
		Score:                 2,
		Label:                 "test",
		DescriptionParagraphs: []string{"p1", "p2"},
	}}

	srvs.On("CalculateResult", mock.Anything).Return(res)

	RegisterHandlers(router, srvs)

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
