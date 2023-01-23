package result

import (
	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
)

type _Repository interface {
	Update([]entity.Result) error
	GetResults() ([]entity.Result, error)
	GetPoint(key string, value string) ([2]int32, error)
	GetResultByFlag(flag int32) (Result, error)
}

type mockRepository struct {
	items []entity.Result
	point [2]int32
	res   Result
}

func (r mockRepository) Update(results []entity.Result) error {
	return nil
}

func (r mockRepository) GetResults() ([]entity.Result, error) {
	return r.items, nil
}

func (r mockRepository) GetPoint(key string, value string) ([2]int32, error) {
	return r.point, nil
}

func (r mockRepository) GetResultByFlag(flag int32) (Result, error) {
	return r.res, nil
}
