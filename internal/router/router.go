package router

import (
	"database/sql"
	"net/http"

	"github.com/coworker-match-api/internal/infra/databases"
	"github.com/coworker-match-api/internal/interfaces/controllers"
	"github.com/coworker-match-api/internal/routers/api"
	"github.com/coworker-match-api/internal/routers/middleware"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) http.Handler {
	h := api.NewHandler(db)
	userRepo := databases.NewSQLUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)

	router := mux.NewRouter()

	// 非認証エンドポイント
	router.HandleFunc("/api/ping", h.PingHandler).Methods("GET")

	// 認証エンドポイントのグループ化
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Auth)

	authRouter.HandleFunc("/users", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/users/{user_id}", userController.GetUser).Methods("GET")
	authRouter.HandleFunc("/users/{user_id}", userController.UpdateUser).Methods("PUT")
	authRouter.HandleFunc("/users/exists", h.UserExistHandler)
	authRouter.HandleFunc("/hobbies", h.HobbyHandler)
	authRouter.HandleFunc("/user_hobbies", h.CreateUserHobbyHandler)
	authRouter.HandleFunc("/user_hobbies/{user_id}", h.UserHobbyHandler)
	authRouter.HandleFunc("/matching_questions", h.MatchingQuestionHandler)
	authRouter.HandleFunc("/matchings", h.MatchingHandler)
	authRouter.HandleFunc("/matchings/{matching_id}", h.MatchingUserHandler)
	authRouter.HandleFunc("/messages", h.MessageHandler).Methods("POST")
	authRouter.HandleFunc("/messages/{message_id}", h.MessageHandler)
	authRouter.HandleFunc("/question_cards", h.QuestionCardHandler)

	return middleware.CORS(router)
}
