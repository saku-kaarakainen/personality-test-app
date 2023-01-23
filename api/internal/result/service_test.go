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

// TODO: Test that file storing fails?
func TestResultStoreFile(t *testing.T) {
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

	if err := s.StoreFile("some_dir"); err != nil {
		assert.FailNow(t, fmt.Sprintf(
			"saving failed with error: %s",
			err.Error(),
		))
	}

	r.AssertCalled(t, "Update", q)
}

// TODO: Should test:
//  1. points are always added to the score,
//     no matter what is gotten from repo
//  2. if error occurs, it will be passed on
//  3. params other than q[%s] will be skipped
func TestCalculateResult(t *testing.T) {
	var err error = nil
	kvps := map[string][]string{
		"q[0]": {"val0", "val1"},
		"q[1]": {"val2", "val3"},
		"q[2]": {"val4", "val5"},
		"q[3]": {"val6", "val7"},
		"q[4]": {"val8", "val9"},
	}

	res := Result{entity.Result{
		Id:                    "1",
		Score:                 3,
		Label:                 "label",
		DescriptionParagraphs: []string{"p1", "p2"},
	}}

	r := new(mockRepository)
	r.On("GetPoint", mock.Anything, mock.Anything).Return([2]int32{1, -1}, err)
	r.On("GetResultByFlag", mock.Anything).Return(res, err)

	s := NewService(r, test.MockLoader{})

	actual, _ := s.CalculateResult(kvps)
	assert.Equal(t, res, actual)
}

func TestAddPointToScore(t *testing.T) {
	var tests = []struct {
		score     [2]int32
		point     [2]int32
		scoreWant [2]int32
	}{
		// test increment
		{[2]int32{0, 0}, [2]int32{1, 1}, [2]int32{1, 1}},

		// decrease
		{[2]int32{-3, -3}, [2]int32{-3, -3}, [2]int32{-6, -6}},

		// diagonal change
		{[2]int32{-3, -3}, [2]int32{6, 6}, [2]int32{3, 3}},
		{[2]int32{2, 2}, [2]int32{-4, -4}, [2]int32{-2, -2}},

		// zero change
		{[2]int32{7, 7}, [2]int32{0, 0}, [2]int32{7, 7}},
		{[2]int32{0, 0}, [2]int32{0, 0}, [2]int32{0, 0}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("addPointToScor (%d): %v", i, tt)

		t.Run(testname, func(t *testing.T) {
			addPointToScore(&tt.score, &tt.point)
			assert.Equal(t, tt.scoreWant, tt.score)
		})
	}
}

// TODO: Test out of range - cases
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
