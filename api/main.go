package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/config"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/db"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/question"
	"github.com/saku-kaarakainen/personality-test-app/api/internal/result"
)

func main() {
	cfg, err := config.Load("./config/config.toml")
	if err != nil {
		panic(err)
	}

	// setup router
	ctx := context.Background()
	router := gin.Default()
	useCors(router, cfg)

	// setup database
	db := db.NewRedisDb(ctx, cfg)
	pong, err := db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("ping: ", pong)

	qsrvs := question.NewService(question.NewRepository(db))
	if err = qsrvs.StoreFile("./config/questions.json"); err != nil {
		panic(err)
	}

	rsrvs := result.NewService(result.NewRepository(db))
	if err = rsrvs.StoreFile("./config/results.json"); err != nil {
		panic(err)
	}

	// attaching handlers to router
	question.RegisterHandlers(router, qsrvs)
	result.RegisterHandlers(router, rsrvs)

	router.Run(cfg.Api.Addr)
}

func useCors(router *gin.Engine, cfg config.Config) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.Api.AllowOrigins
	router.Use(cors.New(corsConfig))
}
