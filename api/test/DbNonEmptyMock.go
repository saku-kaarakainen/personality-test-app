package test

import (
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/stretchr/testify/mock"
)

type DbNonEmptyMock struct{ mock.Mock }

func (m *DbNonEmptyMock) Ping() {

}

func (m *DbNonEmptyMock) Populate() {

}

func (m *DbNonEmptyMock) GetGuestions() ([]db.Question, error) {
	return []db.Question{{
		Id:   "test_id_1",
		Text: "test_text_1",
		Answers: []db.Answer{{
			Id:    "test_id_2",
			Label: "test_label_2",
		}},
	}}, nil
}

func (m *DbNonEmptyMock) GetPoint(key string, value string) ([2]int32, error) {
	return [2]int32{0, 0}, nil
}

func (m *DbNonEmptyMock) GetResult(score [2]int32) (db.Result, error) {
	return db.Result{}, nil
}
