package main

import (
	"context"

	"github.com/gin-gonic/gin"
	// uses redis7
	"github.com/saku-kaarakainen/personality-test-app/api/config"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/server"
)

func main() {
	ctx := context.Background()
	router := gin.Default()
	cfg := config.LoadConfigFromFile("./config/config.toml")

	// database
	db := db.NewDb(cfg, ctx)
	db.Ping()
	db.Populate()

	// server
	srv := server.NewServer(db, router, cfg)
	srv.SetupCors()
	srv.SetRoutes()
	srv.Run()
}
