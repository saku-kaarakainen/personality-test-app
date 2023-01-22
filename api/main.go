package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/db"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/question"
)

// TODO: Use real
type Config struct {
	Addr, Password string
	Db             int
}

func main() {
	// main should:
	// 		init config
	//		tell which db to use
	cfg := Config{}
	ctx := context.Background()
	db := db.NewRedisDb(ctx, cfg.Addr, cfg.Password, cfg.Db)
	logger := log.Default()

	//		run server at port
	//			-> note port value specified at here
	//		handlers
	buildHandler(db, *logger) // TODO: attach the handler to server
	//			-> here. /questions -> points questionService.GetQuestions etc
	//
}

// accepts:
//   - db
//   - config
//   - return handler
func buildHandler(db *db.RedisDb, logger log.Logger) {
	router := gin.Default()

	// TODO: Specify database

	// Questions
	repo := question.NewRepository(db, logger)
	service := question.NewService(repo, logger)
	question.RegisterHandlers(router, service, logger)

	// Answers?

	// Results
}
