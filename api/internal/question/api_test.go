package question

import (
	"net/http"
	"testing"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/test"
)

func TestAPI(t *testing.T) {
	router := test.MockGinRouter()
	srvs := new(mockService)

	qsts := []Question{
		{entity.Question{
			Id:          "1",
			Text:        "text",
			Description: "desc",
			Answers: []entity.Answer{{
				Id:    "1_1",
				Score: [2]int{0, 0},
				Label: "label",
			}},
		}},
	}

	srvs.On("GetQuestions").Return(qsts)

	RegisterHandlers(router, srvs)

	tests := []test.APITestCase{
		{
			Name:         "question: get questions",
			Method:       "GET",
			URL:          "/questions",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusOK,
			WantResponse: `[{"id":"1","question_text":"text","question_description":"desc","answers":[{"id":"1_1","score":[0,0],"answer_label":"label"}]}]`,
		},
		{
			Name:         "question: unknown route",
			Method:       "GET",
			URL:          "/questions/something",
			Body:         "",
			Header:       nil,
			WantStatus:   http.StatusNotFound,
			WantResponse: "",
		},
		{
			Name:         "question: wrong verb",
			Method:       "POST",
			URL:          "/questions",
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
