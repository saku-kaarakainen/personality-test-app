package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/config"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/db"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/question"
)

func main() {
	// main should:
	// 		init config
	//		tell which db to use
	cfg, err := config.Load("./config/config.toml")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// setup database
	db := db.NewRedisDb(ctx, cfg)
	pong, err := db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("ping: ", pong)

	qsrvs := question.NewService(question.NewRepository(db))
	err = qsrvs.StoreFile("./config/questions.json")
	if err != nil {
		panic(err)
	}

	//rsrvs := result.NewService(result.NewRepository(db, *logger), *logger)
	log.Println("setting router")
	router := gin.Default()
	useCors(router, cfg)

	// Questions

	question.RegisterHandlers(router, qsrvs)

	// Answers?

	// TODO: Results
	// result.RegisterHandlers(
	// 	router,
	// 	result.NewService(
	// 		result.NewRepository(db, logger),
	// 		logger,
	// 	),
	// 	logger,
	// )

	router.Run(cfg.Api.Addr)
}

func useCors(router *gin.Engine, cfg config.Config) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Api.AllowOrigins
	router.Use(cors.New(corsConfig))
}
