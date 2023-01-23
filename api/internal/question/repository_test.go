package question

import (
	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (r *mockRepository) Update(questions []entity.Question) error {
	args := r.Called(questions)
	return args.Error(0)
}

func (r *mockRepository) GetQuestions() ([]entity.Question, error) {
	args := r.Called()
	q := args.Get(0)
	questions := q.([]entity.Question)

	return questions, args.Error(1)
}
