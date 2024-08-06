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
	hobbyRepo := databases.NewSQLHobbyRepository(db)
	hobbyUsecase := usecases.NewHobbyUsecase(hobbyRepo)
	hobbyController := controllers.NewHobbyController(hobbyUsecase)

	router := mux.NewRouter()

	// 非認証エンドポイント
	router.HandleFunc("/api/ping", h.PingHandler).Methods("GET")

	// 認証エンドポイントのグループ化
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Auth)

	authRouter.HandleFunc("/users/exists", userController.IsUserExist).Methods("GET")
	authRouter.HandleFunc("/users", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/users/{userId}", userController.GetUser).Methods("GET")
	authRouter.HandleFunc("/users/{userId}", userController.UpdateUser).Methods("PUT")

	authRouter.HandleFunc("/hobbies", hobbyController.GetAllHobby).Methods("GET")

	authRouter.HandleFunc("/user_hobbies", h.CreateUserHobbyHandler)
	authRouter.HandleFunc("/user_hobbies/{userId}", h.UserHobbyHandler)
	authRouter.HandleFunc("/matching_questions", h.MatchingQuestionHandler)
	authRouter.HandleFunc("/matchings", h.MatchingHandler)
	authRouter.HandleFunc("/matchings/{matchingId}", h.MatchingUserHandler)
	authRouter.HandleFunc("/messages", h.MessageHandler).Methods("POST")
	authRouter.HandleFunc("/messages/{messageId}", h.MessageHandler)
	authRouter.HandleFunc("/question_cards", h.QuestionCardHandler)

	return middleware.CORS(router)
}
