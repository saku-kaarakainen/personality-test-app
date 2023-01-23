package question

import (
	"encoding/json"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/utils"
)

type Service interface {
	StoreFile(filename string) error
	GetQuestions() ([]Question, error)
}

type Question struct{ entity.Question }

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
	var questions []entity.Question
	json.Unmarshal(byteValue, &questions)

	// 3. store file
	// Note: This is redis database, so the value will be inserted if it does not exist.
	if err := s.repo.Update(questions); err != nil {
		return err
	}

	return nil
}

func (s service) GetQuestions() ([]Question, error) {
	data, err := s.repo.GetQuestions()
	if err != nil {
		return nil, err
	}

	var questions []Question
	for _, q := range data {
		questions = append(questions, Question{q})
	}

	return questions, nil
}
