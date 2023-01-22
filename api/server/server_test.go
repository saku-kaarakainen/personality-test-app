package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/saku-kaarakainen/personality-test-app/api/config"
	"github.com/saku-kaarakainen/personality-test-app/api/test"
)

var (
	fakeDb = new(test.DbNonEmptyMock)
	router = gin.Default()
	server = NewServer(fakeDb, router, &config.Config{})
	w      *httptest.ResponseRecorder
)

func init() {
	// add routes, because those are being tested.
	server.SetRoutes()
	w = httptest.NewRecorder()
}

// Test route /ping exists and responds HTTP 200 with "pong"
func TestPingRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// Test route /questions exists and responds HTTP 200 returns values using db.module
func TestGetQuestionsRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/questions", nil)
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[{"id":"test_id_1","question_text":"test_text_1","question_description":"","answers":[{"id":"test_id_2","score":[0,0],"answer_label":"test_label_2"}]}]`, w.Body.String())
}
