package routes

import (
	"github.com/gin-gonic/gin"
)

func Get_questions(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}
