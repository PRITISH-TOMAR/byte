package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"by_te/internal/config"
	"by_te/internal/infrastructure/logger"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Logger.Fatalf("db connection error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.DBTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		logger.Logger.Fatalf("db ping failed: %v", err)
	}

	logger.Logger.Println("connected to mysql")

	return db
}