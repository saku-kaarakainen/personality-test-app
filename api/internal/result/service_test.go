package result

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestResultStoreFile(t *testing.T) {
	// internal/question is the current dir
	filename := "./../../config/results.json"
	var err error = nil
	q := []entity.Result{{
		Id:    "1",
		Label: "label",
		Score: 3,
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

func TestConvertScoreToFlag(t *testing.T) {

	var tests = []struct {
		score        [2]int32
		expectedFlag int32
	}{
		{[2]int32{-2, -2}, 0},
		{[2]int32{2, -2}, 1},
		{[2]int32{-2, 2}, 2},
		{[2]int32{2, 2}, 3},
		// There is no error handling for indexing, so this might fail in special cases, like if %s is at end.
	}

	for i, tt := range tests {
		testname := fmt.Sprintf(
			"%d (for score:'%v', expected:'%d')",
			i, tt.score, tt.expectedFlag)

		t.Run(testname, func(t *testing.T) {
			actualFlag := convertScoreToFlag(tt.score)
			assert.Equal(t, tt.expectedFlag, actualFlag)
		})
	}
}

type mockService struct {
	mock.Mock
}

func (s *mockService) StoreFile(filename string) error {
	return nil
}

func (s *mockService) CalculateResult(kvps map[string][]string) (Result, error) {
	args := s.Called(kvps)
	return args.Get(0).(Result), nil
}
