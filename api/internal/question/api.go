package question

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, service Service) {
	h := handler{service: service}

	// handlers relatec to /questions
	router.GET("/questions", h.getQuestions)
}

type handler struct {
	service Service
}

func (h handler) getQuestions(ctx *gin.Context) {
	log.Println("/questions")

	data, err := h.service.GetQuestions()
	if err != nil {
		log.Println("Error getting value:", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "An error occurred",
		})
		return
	}

	log.Println("Response: HTTP 200")
	log.Println(data)
	ctx.JSON(200, data)
}
