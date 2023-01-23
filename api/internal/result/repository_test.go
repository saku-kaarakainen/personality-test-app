package result

import (
	"github.com/saku-kaarakainen/personality-test-app/api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (r *mockRepository) Update(results []entity.Result) error {
	args := r.Called(results)
	return args.Error(0)
}

func (r mockRepository) GetResults() ([]entity.Result, error) {
	args := r.Called()
	res := args.Get(0)
	results := res.([]entity.Result)

	return results, args.Error(1)
}

func (r mockRepository) GetPoint(key string, value string) ([2]int32, error) {
	args := r.Called(key, value)
	p := args.Get(0)
	point := p.([2]int32)

	return point, args.Error(1)
}

func (r mockRepository) GetResultByFlag(flag int32) (Result, error) {
	args := r.Called(flag)
	res := args.Get(0)
	result := res.(Result)

	return result, args.Error(1)
}
