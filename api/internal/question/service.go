package question

import (
	"log"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
)

type Service interface {
	GetQuestions() ([]Question, error)
}

type Question struct{ entity.Question }

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
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
