package routes

import (
	"encoding/json"
	"fmt"

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
	jsonBlob, err := db.Get("test")
	var data map[string]interface{}
	json.Unmarshal(jsonBlob.([]byte), &data)

	// fmt.Println("We found: " + string(obj))

	// //fmt.Println("No errors! So we are going to pass json file: " + jsonFile)
	// ctx.JSON(200, string(obj))

	// p := Person{Name: "John Doe", Age: 30}

	// jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling:", err)
	}

	// test := gin.H{}

	ctx.JSON(200, data)
}
