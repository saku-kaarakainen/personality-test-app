package test

import (
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/stretchr/testify/mock"
)

type DbEmptyMock struct{ mock.Mock }

func (m *DbEmptyMock) Ping() {

}

func (m *DbEmptyMock) Populate() {

}

func (m *DbEmptyMock) GetGuestions() ([]db.Question, error) {
	return nil, nil
}

func (m *DbEmptyMock) GetPoint(key string, value string) ([2]int32, error) {
	return [2]int32{0, 0}, nil
}

func (m *DbEmptyMock) GetResult(score [2]int32) (db.Result, error) {
	return db.Result{}, nil
}
