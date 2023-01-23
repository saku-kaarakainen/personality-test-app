package test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func MockGinRouter() *gin.Engine {
	_, router := gin.CreateTestContext(httptest.NewRecorder())

	return router
}

/*
TODO: Remove
// MockRouter creates a routing.Router for testing APIs.
func MockRouterOld(logger log.Logger) *routing.Router {
	router := routing.New()
	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)
	return router
}
*/
