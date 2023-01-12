package main

import (
	"log"

	"github.com/gin-gonic/gin"

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
	log.Println("Start debugging at http://localhost:8080/questions")
}
