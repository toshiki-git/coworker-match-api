package routers

import (
	"database/sql"
	"net/http"

	"github.com/coworker-match-api/internal/routers/api"
)

func InitRouter(db *sql.DB) http.Handler {
	h := api.NewHandler(db)
	mux := http.NewServeMux()

	mux.HandleFunc("/api/ping", h.PingHandler)
	mux.HandleFunc("/api/users", h.UsersHandler)

	return mux
}
