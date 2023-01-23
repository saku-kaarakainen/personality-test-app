package result

import (
	"encoding/json"
	"log"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/utils"
	myutils "github.com/saku-kaarakainen/personality-test-app/api/utils"
)

type Service interface {
	StoreFile(filename string) error
	CalculateResult(kvps map[string][]string) (Result, error)
}

type Result struct{ entity.Result }

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

// Stores loaded file in database
//
// The logic is stored at service level as one func, because
// this logic is requisite for server to operate.
func (s service) StoreFile(filename string) error {
	// 1. load the file
	byteValue, err := myutils.LoadFile(filename)
	if err != nil {
		log.Println("load Results file failed")
		return err
	}

	// 2. cast to correct type
	var Results []entity.Result
	json.Unmarshal(byteValue, &Results)

	log.Println("loaded Results file ")

	// 3. store file
	// Note: This is redis database, so the value will be inserted if it does not exist.
	if err := s.repo.Update(Results); err != nil {
		log.Println("updating Results failed")
		return err
	}

	log.Println("file stored")
	return nil
}

func (s service) CalculateResult(kvps map[string][]string) (Result, error) {
	log.Println("SERVICE CALCULATE RESULT")
	score := [2]int32{0, 0}

	log.Println("kvps: ", kvps)

	// Note: "Business logic"
	for raw_key, value_array := range kvps {
		// Get the index from the url parameter
		key, err := utils.Unformat("q[%s]", raw_key)
		if err != nil {
			continue
		}

		// Get the key and value from the param
		value := value_array[0]

		log.Printf("key '%s', value '%s'\n", key, value)

		point, err := s.repo.GetPoint(key, value)
		if err != nil {
			return Result{}, nil
		}

		log.Println("found point: ", point)
		score[0] += point[0]
		score[1] += point[1]

		log.Println("score is now: ", score)
	}

	log.Println("convert score '%v' to flag.", score)
	flag := convertScoreToFlag(score)
	log.Println("generated flag: ", flag)

	result, err := s.repo.GetResultByFlag(flag)
	log.Println("found result: ", result)

	if err != nil {
		return Result{}, err
	}

	return result, nil
}

// Converts two dimensional score into binary flag.
// Checks if the point of the score is above zero,
// That is interpret as true.
// So, only the sign of the integer matters, not the value.
// Please check the related unit test for how to use it
//
// If the CPU usage turnes out the be the performance bottle neck,
// you can optimize the code furthermode with using this one liner instead:
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
