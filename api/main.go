package main

import (
	"context"

	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis/v8" // uses redis7
	"github.com/nitishm/go-rejson/v4"

	"github.com/saku-kaarakainen/personality-test-app/api/config"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/routes"
)

func setupRouter(db db.IDb) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		// It is more approariate to put the func into it's own file, 'routes/ping.go'.
		// However this goes easily into very big rabbit hole with better framework, or better use of it.
		// Right now it's better to keep the code simpler and leave this as-is.
		ctx.String(200, "pong")
	})
	router.GET("/questions", func(ctx *gin.Context) {
		routes.Get_questions(ctx, db)
	})
	return router
}

func setupDatabase() *db.Db {
	var (
		ctx = context.Background()
		cli = goredis.NewClient(&goredis.Options{
			Addr:     config.Db.Addr,
			Password: config.Db.Pw,
			DB:       config.Db.SelectedDb,
		})
		rh = rejson.NewReJSONHandler()
	)

	rh.SetGoRedisClient(cli)
	database := db.NewDb(ctx, cli, rh)
	database.Ping()
	database.Populate()
	return database
}

func main() {
	database := setupDatabase()
	router := setupRouter(database)

	router.Run(config.Api.Addr)
}
