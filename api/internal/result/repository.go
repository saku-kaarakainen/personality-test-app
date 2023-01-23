package result

import (
	"encoding/json"
	"fmt"

	"github.com/saku-kaarakainen/personality-test-app/api/internal/db"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
)

type Repository interface {
	Update([]entity.Result) error
	GetResults() ([]entity.Result, error)
	GetPoint(key string, value string) ([2]int32, error)
	GetResultByFlag(flag int32) (Result, error)
}

type repository struct {
	db db.NoSqlDb
}

func NewRepository(db db.NoSqlDb) Repository {
	return repository{db: db}
}

func (r repository) Update(Results []entity.Result) error {
	return r.db.Update("results", Results)
}

func (r repository) GetResults() ([]entity.Result, error) {
	jsonBlob, err := r.db.Get("results", ".")
	if err != nil {
		return nil, err
	}

	var data []entity.Result
	json.Unmarshal(jsonBlob.([]byte), &data)

	return data, nil
}

func (r repository) GetPoint(key string, value string) ([2]int32, error) {
	path := fmt.Sprintf("$.[?(@.id==\"%s\")].answers[?(@.id==\"%s\")].score", key, value)
	jsonBlob, err := r.db.Get("questions", path)
	if err != nil {
		return [2]int32{0, 0}, err
	}

	var data [1][2]int32
	json.Unmarshal(jsonBlob.([]byte), &data)

	return data[0], nil
}

func (r repository) GetResultByFlag(flag int32) (Result, error) {
	path := fmt.Sprintf("$.[?(@.score==%d)]", flag)
	jsonBlob, err := r.db.Get("results", path)
	if err != nil {
		return Result{}, err
	}

	var data []Result
	json.Unmarshal(jsonBlob.([]byte), &data)

	if len(data) == 0 {
		return Result{}, fmt.Errorf("result not found with score '%d'", flag)
	}

	return data[0], nil
}
