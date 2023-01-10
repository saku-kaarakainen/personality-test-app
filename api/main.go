package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9" // uses redis7

	"github.com/saku-kaarakainen/personality-test-app/api/routes"
)

var ctx = context.Background()

func setupRedis() {
	rdb := redis.NewClient(&redis.Options{
		// TODO: Add username and password, also .env
		Addr:     "personality-test-db:6379",
		Password: "", // no password set
		DB:       0,  // use default db
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println("Redis Ping:")
	fmt.Println(pong, err)

	if err != nil {
		panic(err)
	}
}

func setupServer() {
	server := gin.Default()
	server.GET("/questions", routes.Get_questions)
	server.Run(":8080")
}

func main() {
	setupRedis()
	setupServer()
}
