package routers

import (
	"database/sql"
	"net/http"

	"github.com/coworker-match-api/internal/routers/api"
	"github.com/coworker-match-api/internal/routers/middleware"
)

func InitRouter(db *sql.DB) http.Handler {
	h := api.NewHandler(db)
	mux := http.NewServeMux()

	// 非認証エンドポイント
	mux.HandleFunc("/api/ping", h.PingHandler)

	// 認証エンドポイント
	mux.Handle("/api/users", middleware.Auth(http.HandlerFunc(h.CreateUserHandler)))
	mux.Handle("/api/users/{user_id}", middleware.Auth(http.HandlerFunc(h.UserHandler)))
	mux.Handle("/api/hobbies", middleware.Auth(http.HandlerFunc(h.HobbyHandler)))

	return middleware.CORS(mux)
}
