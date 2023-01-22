package services

import (
	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
)

type Question struct {
	ctx *gin.Context
	db  db.IDb
}

func NewQuestion(
	ctx *gin.Context,
	db db.IDb,
) Question {
	return Question{
		ctx: ctx,
		db:  db,
	}
}

func (q *Question) GetQuestions() ([]db.Question, error) {
	data, err := q.db.GetGuestions()
	if err != nil {
		return []db.Question{}, nil
	}

	return data, nil
}
