package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/utils"
)

func Get_Result_Calculate(ctx *gin.Context, db db.IDb) {
	log.Println("Entered method Get Result Calculate")
	score := [2]int32{0, 0}

	log.Println("We are iterating keys:")
	log.Println(ctx.Request.URL.Query())

	for raw_key, value_array := range ctx.Request.URL.Query() {
		format := "q[%s]"
		// Get the index from the url parameter
		//raw_key := "q[3]"
		key, err := utils.Unformat(format, raw_key)
		if err != nil {
			log.Printf("Skipped key '%s'.\n", key)
			continue
		}

		log.Println(key)
		// expected output: 3

		// Get the key and value from the param
		value := value_array[0]
		log.Println("Found value:")
		log.Println(value)
		log.Println()

		point, err := db.GetPoint(key, value)
		log.Println("got point:")
		log.Println(point)

		if err != nil {
			log.Println("Error getting value:", err)
			AbortMsg(500, err, ctx)
			return
		}

		log.Println("updating score")
		// add points to the score
		score[0] += point[0]
		score[1] += point[1]

		log.Println("Score is:")
		log.Println(score)
	}

	log.Println("Get results")
	result, err := db.GetResult(score)
	if err != nil {
		log.Println("Error getting value:", err)
		AbortMsg(500, err, ctx)
		return
	}

	log.Println("result:")
	log.Println(result)
	ctx.String(420, "An error occurred")
}
