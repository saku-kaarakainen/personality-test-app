package question

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, service Service) {
	h := handler{service: service}

	router.GET("/questions", h.getQuestions)
}

type handler struct {
	service Service
}

func (h handler) getQuestions(ctx *gin.Context) {
	data, err := h.service.GetQuestions()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "An error occurred",
		})
		return
	}

	ctx.JSON(200, data)
}
