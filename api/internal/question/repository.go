package question

import (
	"encoding/json"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/db"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
)

type Repository interface {
	Update([]entity.Question) error
	GetQuestions() ([]entity.Question, error)
}

type repository struct {
	db db.NoSqlDb
}

func NewRepository(db db.NoSqlDb) Repository {
	return repository{db: db}
}

func (r repository) Update(questions []entity.Question) error {
	return r.db.Update("questions", questions)
}

func (r repository) GetQuestions() ([]entity.Question, error) {
	jsonBlob, err := r.db.Get("questions", ".")
	if err != nil {
		return nil, err
	}

	var data []entity.Question
	json.Unmarshal(jsonBlob.([]byte), &data)

	return data, nil
}

// TODO: Set questions
