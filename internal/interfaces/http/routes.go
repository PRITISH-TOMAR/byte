package http

import (
	"database/sql"
	"net/http"

	goredis "github.com/redis/go-redis/v9"
)

func RegisterRoutes(db *sql.DB, rdb *goredis.Client) {
	http.HandleFunc("/health", healthHandler(db, rdb))
}
