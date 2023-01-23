package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, service Service) {
	h := handler{service: service}

	// handlers relatec to /questions
	router.GET("/result/calculate", h.getResult)
}

type handler struct {
	service Service
}

func (h handler) getResult(ctx *gin.Context) {
	data, err := h.service.CalculateResult(ctx.Request.URL.Query())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "An error occurred",
		})
		return
	}

	ctx.JSON(200, data)
}
