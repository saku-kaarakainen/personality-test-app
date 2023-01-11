package main

import (
	"github.com/gin-gonic/gin"

	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	router.GET("/questions", routes.Get_questions)
	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
	db.LoadModule()
}
