package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis/v8" // uses redis7
	"github.com/nitishm/go-rejson/v4"

	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/routes"
)

func setupRouter(db db.IDb) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
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
			// TODO: Add username and password, also .env
			Addr:     "personality-test-db:6379",
			Password: "", // no password set
			DB:       0,  // use default db
			//Network:     "db-network",
			//DialTimeout: 60,
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
	log.Println("Main.go. Setup database")
	database := setupDatabase()

	log.Println("Database is ok! setup router")
	router := setupRouter(database)

	log.Println("Router is done. Let's run it!")
	router.Run(":8080")
	log.Println("Start debugging at http://localhost:8080/questions")
}
