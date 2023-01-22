package services

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/utils"
)

type Result struct {
	ctx *gin.Context
	db  db.IDb
}

func NewResult(
	ctx *gin.Context,
	db db.IDb,
) Result {
	return Result{
		ctx: ctx,
		db:  db,
	}
}

func (r *Result) CalculateResult() {
	score := [2]int32{0, 0}

	// Note: "Business logic"
	for raw_key, value_array := range r.ctx.Request.URL.Query() {
		// Get the index from the url parameter
		key, err := utils.Unformat("q[%s]", raw_key)
		if err != nil {
			log.Printf("Skipped key '%s'.\n", key)
			continue
		}

		// Get the key and value from the param
		value := value_array[0]
		point, err := r.db.GetPoint(key, value)

		if err != nil {
			log.Println("Error getting value:", err)
			AbortMsg(500, err, ctx)
			return
		}

		log.Println("updating score")

		// add points to the score
		// Note: "businesss logic"
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

	ctx.JSON(200, result)
}
