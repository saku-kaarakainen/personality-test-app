package main

import (
	"github.com/gin-gonic/gin"

	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/routes"
)

func setupServer() {
	server := gin.Default()
	server.GET("/questions", routes.Get_questions)
	server.Run(":8080")
}

func main() {
	db.LoadModule()
	setupServer()
}
