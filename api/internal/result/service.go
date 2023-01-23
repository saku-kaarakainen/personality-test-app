package result

import (
	"encoding/json"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/utils"
)

type Service interface {
	StoreFile(filename string) error
	CalculateResult(kvps map[string][]string) (Result, error)
}

type Result struct{ entity.Result }

type service struct {
	repo   Repository
	loader utils.Loader
}

func NewService(repo Repository, loader utils.Loader) Service {
	return service{
		repo:   repo,
		loader: loader,
	}
}

// Stores loaded file in database
//
// The logic is stored at service level as one func, because
// this logic is requisite for server to operate.
func (s service) StoreFile(filename string) error {
	// 1. load the file
	byteValue, err := s.loader.LoadFile(filename)
	if err != nil {
		return err
	}

	// 2. cast to correct type
	var Results []entity.Result
	json.Unmarshal(byteValue, &Results)

	// 3. store file
	// Note: This is redis database, so the value will be inserted if it does not exist.
	if err := s.repo.Update(Results); err != nil {
		return err
	}

	return nil
}

func (s service) CalculateResult(kvps map[string][]string) (Result, error) {
	score := [2]int32{0, 0}

	for raw_key, value_array := range kvps {
		// Get the index from the url parameter
		key, err := utils.Unformat("q[%s]", raw_key)
		if err != nil {
			continue
		}

		// Get the key and value from the param
		value := value_array[0]
		point, err := s.repo.GetPoint(key, value)
		if err != nil {
			return Result{}, nil
		}

		// Add points to score
		score[0] += point[0]
		score[1] += point[1]
	}

	flag := convertScoreToFlag(score)
	result, err := s.repo.GetResultByFlag(flag)
	if err != nil {
		return Result{}, err
	}

	return result, nil
}

// Converts two dimensional score into a flag.
// Flag can be a number between 0 - 3.
// Checks if the number of a score is positive
// and converts positive number into true
// and negative into false.
//
//	+----------+----------+-----+
//	| score[0] | score[1] | ret |
//	+----------+----------+-----+
//	| neg int  | neg int  | 0   |
//	| neg int  | pos int  | 1   |
//	| pos int  | neg int  | 2   |
//	| pos int  | pos int  | 3   |
//	+----------+----------+-----+
//
// If the CPU usage turnes out the be the performance bottle neck,
// you can optimize the code furthermode by using this one liner instead:
//
//	return (int32((-score[0]>>31)&1) | (int32((-score[1]>>31)&1) << 1))
func convertScoreToFlag(score [2]int32) int32 {
	x_flag := score[0] >= 0
	y_flag := score[1] >= 0

	flags := int32(0)

	if x_flag {
		flags |= 1 << 0
	}

	if y_flag {
		flags |= 1 << 1
	}

	return flags
}
