package router

import (
	"database/sql"
	"net/http"

	"github.com/coworker-match-api/internal/middleware"
	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) http.Handler {

	router := mux.NewRouter()

	// 非認証エンドポイント
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong-pong"))
	}).Methods("GET")

	// 認証エンドポイントのグループ化
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Auth)

	NewUserRouter(db, authRouter)
	NewUserHobbyRouter(db, authRouter)
	NewHobbyRouter(db, authRouter)
	NewQuestionCardRouter(db, authRouter)
	NewMatchingQuestionRouter(db, authRouter)
	NewMatchingRouter(db, authRouter)
	NewMessageRouter(db, authRouter)

	return middleware.CORS(router)
}
