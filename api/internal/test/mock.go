package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func MockGinRouter() *gin.Engine {
	_, router := gin.CreateTestContext(httptest.NewRecorder())

	return router
}

type MockLoader struct {
	RetBytes []byte
}

func (l MockLoader) LoadFile(filename string) ([]byte, error) {
	bytes := l.RetBytes
	return bytes, nil
}
