package main

import (
	"net/http"

	"by_te/internal/config"
	"by_te/internal/infrastructure/cache/redis"
	"by_te/internal/infrastructure/database/mysql"
	"by_te/internal/infrastructure/logger"
	httpInterface "by_te/internal/interfaces/http"
)

func main() {
	// Init logger
	logger.Init()
	logger.Logger.Println("starting application")

	// Load config
	cfg := config.LoadConfig()

	// Init DB
	db := mysql.NewMySQL(cfg)
	defer db.Close()

	// Init Redis
	rdb := redis.NewRedis(cfg)
	defer rdb.Close()

	// Register routes
	httpInterface.RegisterRoutes()

	// Start server
	logger.Logger.Printf("server running on :%s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		logger.Logger.Fatalf("server failed: %v", err)
	}
}