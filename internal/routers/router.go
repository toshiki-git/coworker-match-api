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
	mux.HandleFunc("/api/users", h.CreateUserHandler)
	mux.HandleFunc("/api/users/{user_id}", h.UserHandler)
	mux.HandleFunc("/api/hobbies", h.HobbyHandler)

	return mux
}
