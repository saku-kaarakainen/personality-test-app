package question

import "github.com/saku-kaarakainen/personality-test-app/api/internal/entity"

type mockRepository struct {
	items []entity.Question
}

func (r mockRepository) Update(questions []entity.Question) error {
	return nil
}

func (r mockRepository) GetQuestions() ([]entity.Question, error) {
	return r.items, nil
}
