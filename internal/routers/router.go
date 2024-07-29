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

	// 認証エンドポイントのグループ化
	authMux := http.NewServeMux()
	authMux.HandleFunc("/api/users", h.CreateUserHandler)
	authMux.HandleFunc("/api/users/{user_id}", h.UserHandler)
	authMux.HandleFunc("/api/users/exists", h.UserExistHandler)
	authMux.HandleFunc("/api/hobbies", h.HobbyHandler)
	authMux.HandleFunc("/api/user_hobbies", h.CreateUserHobbyHandler)
	authMux.HandleFunc("/api/user_hobbies/{user_id}", h.UserHobbyHandler)
	authMux.HandleFunc("/api/matching_questions", h.MatchingQuestionHandler)
	authMux.HandleFunc("/api/matchings", h.MatchingHandler)
	authMux.HandleFunc("/api/matchings/{matching_id}", h.MatchingUserHandler)
	authMux.HandleFunc("/api/messages", h.MessageHandler)
	authMux.HandleFunc("/api/messages/{message_id}", h.MessageHandler)
	authMux.HandleFunc("/api/question_cards", h.QuestionCardHandler)

	// 認証ミドルウェア
	mux.Handle("/api/", middleware.Auth(authMux))

	return middleware.CORS(mux)
}
