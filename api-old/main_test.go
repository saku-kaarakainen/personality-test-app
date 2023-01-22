package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DbEmptyMock struct{ mock.Mock }

func (m *DbEmptyMock) Ping()                                {}
func (m *DbEmptyMock) Populate()                            {}
func (m *DbEmptyMock) GetGuestions() ([]db.Question, error) { return nil, nil }
func (m *DbEmptyMock) GetPoint(key string, value string) ([2]int32, error) {
	return [2]int32{0, 0}, nil
}
func (m *DbEmptyMock) GetResult(score [2]int32) (db.Result, error) { return db.Result{}, nil }

// Test route /ping exists and responds HTTP 200 with "pong"
func TestPingRoute(t *testing.T) {
	fakeDb := new(DbEmptyMock)
	router := gin.Default()
	setupRoutes(router, fakeDb)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

type DbNonEmpty struct{ mock.Mock }

func (m *DbNonEmpty) Ping()     {}
func (m *DbNonEmpty) Populate() {}
func (m *DbNonEmpty) GetGuestions() ([]db.Question, error) {
	return []db.Question{{
		Id:   "test_id_1",
		Text: "test_text_1",
		Answers: []db.Answer{{
			Id:    "test_id_2",
			Label: "test_label_2",
		}},
	}}, nil
}
func (m *DbNonEmpty) GetPoint(key string, value string) ([2]int32, error) { return [2]int32{0, 0}, nil }
func (m *DbNonEmpty) GetResult(score [2]int32) (db.Result, error)         { return db.Result{}, nil }

// Test route /questions exists and responds HTTP 200 returns values using db.module
func TestGetQuestionsRoute(t *testing.T) {
	fakeDb := new(DbNonEmpty)
	router := gin.Default()
	setupRoutes(router, fakeDb)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/questions", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[{"id":"test_id_1","question_text":"test_text_1","question_description":"","answers":[{"id":"test_id_2","score":[0,0],"answer_label":"test_label_2"}]}]`, w.Body.String())
}
