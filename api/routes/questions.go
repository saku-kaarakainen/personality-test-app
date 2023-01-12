package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
)

func AbortMsg(code int, err error, c *gin.Context) {
	c.String(code, "An error occurred")
	// A custom error page with HTML templates can be shown by c.HTML()
	c.Error(err)
	c.Abort()
}

type Person struct {
	Name string
	Age  int
}

func Get_questions(ctx *gin.Context) {
	log.Println("get_questions")

	data, err := db.GetGuestions()
	if err != nil {
		log.Println("Error getting value:", err)
		AbortMsg(500, err, ctx)
		return
	}

	log.Println("Return data:")
	log.Println(data)
	ctx.JSON(200, data)
}
