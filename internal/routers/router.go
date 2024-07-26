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
	mux.Handle("/api/users/exists", middleware.Auth(http.HandlerFunc(h.UserExistHandler)))
	mux.Handle("/api/hobbies", middleware.Auth(http.HandlerFunc(h.HobbyHandler)))
	mux.Handle("/api/user_hobbies", middleware.Auth(http.HandlerFunc(h.CreateUserHobbyHandler)))
	mux.Handle("/api/user_hobbies/{user_id}", middleware.Auth(http.HandlerFunc(h.UserHobbyHandler)))
	mux.Handle("/api/matching_questions", middleware.Auth(http.HandlerFunc(h.MatchingQuestionHandler)))
	mux.Handle("/api/matchings", middleware.Auth(http.HandlerFunc(h.MatchingHandler)))
	mux.Handle("/api/messages", middleware.Auth(http.HandlerFunc(h.MessageHandler)))
	mux.Handle("/api/messages/{message_id}", middleware.Auth(http.HandlerFunc(h.MessageHandler)))
	mux.Handle("/api/question_cards", middleware.Auth(http.HandlerFunc(h.QuestionCardHandler)))

	return middleware.CORS(mux)
}
