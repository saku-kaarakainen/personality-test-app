package question

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQuestionStoreFile(t *testing.T) {
	// internal/question is the current dir
	filename := "./../../config/questions.json"
	var err error = nil
	q := []entity.Question{{
		Id:          "1",
		Text:        "text",
		Description: "desc",
	}}

	b, _ := json.Marshal(q)

	r := new(mockRepository)
	r.On("Update", q).Return(err)

	s := NewService(r, test.MockLoader{
		RetBytes: b,
	})
	// Act
	if err := s.StoreFile(filename); err != nil {
		assert.FailNow(t, fmt.Sprintf(
			"saving '%s' failed with error: %s",
			filename,
			err.Error(),
		))
	}

	// assert
	r.AssertCalled(t, "Update", q)
}

func TestGetQuestions(t *testing.T) {
	var err error = nil
	expected := []Question{{entity.Question{
		Id:          "1",
		Text:        "text",
		Description: "desc",
	}}}
	r := new(mockRepository)
	r.On("GetQuestions").Return([]entity.Question{{
		Id:          "1",
		Text:        "text",
		Description: "desc",
	}}, err)

	s := NewService(r, test.MockLoader{})

	actual, _ := s.GetQuestions()

	assert.Equal(t, expected, actual)
}

type mockService struct {
	mock.Mock
}

func (s *mockService) StoreFile(filename string) error {
	return nil
}

func (s *mockService) GetQuestions() ([]Question, error) {
	return s.Called().Get(0).([]Question), nil
}
