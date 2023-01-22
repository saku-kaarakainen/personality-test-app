package question

import (
	"encoding/json"
	"log"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/db"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
)

type Repository interface {
	GetQuestions() ([]entity.Question, error)
}

type repository struct {
	db     *db.RedisDb
	logger log.Logger
}

func NewRepository(db *db.RedisDb, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) GetQuestions() ([]entity.Question, error) {
	jsonBlob, err := r.db.Get("questions", ".")
	if err != nil {
		return nil, err
	}

	var data []entity.Question
	json.Unmarshal(jsonBlob, &data)

	return data, nil
}

// TODO: Set questions
