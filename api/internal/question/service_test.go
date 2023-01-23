package question

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStoreFile(t *testing.T) {
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
	log.Println(" (1) calls:", r.Calls)

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

	sv := s.(service)

	log.Println(" (3) calls:", r.Calls)
	log.Println("r: ", r)
	log.Println("sv.repo:", sv.repo)

	log.Println("assert Update with q: ", q)
	// assert
	r.AssertCalled(t, "Update", q)
}

type mockService struct {
	mock.Mock
}

func (s mockService) StoreFile(filename string) error {
	return nil
}

func (s mockService) GetQuestions() ([]Question, error) {
	return s.Called().Get(0).([]Question), nil
}
