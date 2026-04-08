package http

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type healthStatus struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks"`
}

func healthHandler(db *sql.DB, rdb *goredis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		checks := make(map[string]string)
		overall := "ok"

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			checks["mysql"] = "error: " + err.Error()
			overall = "degraded"
		} else {
			checks["mysql"] = "ok"
		}

		if err := rdb.Ping(ctx).Err(); err != nil {
			checks["redis"] = "error: " + err.Error()
			overall = "degraded"
		} else {
			checks["redis"] = "ok"
		}

		resp := healthStatus{Status: overall, Checks: checks}

		w.Header().Set("Content-Type", "application/json")
		if overall != "ok" {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		json.NewEncoder(w).Encode(resp)
	}
}
