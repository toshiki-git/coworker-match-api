package router

import (
	"database/sql"
	"net/http"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/routers/api"
	"github.com/coworker-match-api/internal/routers/middleware"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) http.Handler {
	h := api.NewHandler(db)
	userRepo := repositories.NewUserRepo(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)
	hobbyRepo := repositories.NewHobbyRepo(db)
	hobbyUsecase := usecases.NewHobbyUsecase(hobbyRepo)
	hobbyController := controllers.NewHobbyController(hobbyUsecase)
	userHobbyRepo := repositories.NewUserHobbyRepo(db)
	userHobbyUsecase := usecases.NewUserHobbyUsecase(userHobbyRepo)
	userHobbyController := controllers.NewUserHobbyController(userHobbyUsecase)

	router := mux.NewRouter()

	// 非認証エンドポイント
	router.HandleFunc("/api/ping", h.PingHandler).Methods("GET")

	// 認証エンドポイントのグループ化
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Auth)

	authRouter.HandleFunc("/users/exists", userController.IsUserExist).Methods("GET")
	authRouter.HandleFunc("/users", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/users/{userId}", userController.GetUserById).Methods("GET")
	authRouter.HandleFunc("/users/{userId}", userController.UpdateUser).Methods("PUT")

	authRouter.HandleFunc("/hobbies", hobbyController.GetAllHobby).Methods("GET")

	authRouter.HandleFunc("/user_hobbies", userHobbyController.CreateUserHobby).Methods("POST")
	authRouter.HandleFunc("/user_hobbies", userHobbyController.UpdateUserHobby).Methods("PUT")
	authRouter.HandleFunc("/user_hobbies/{userId}", userHobbyController.GetAllUserHobby).Methods("GET")

	authRouter.HandleFunc("/matching_questions", h.MatchingQuestionHandler)
	authRouter.HandleFunc("/matchings", h.MatchingHandler)
	authRouter.HandleFunc("/matchings/{matchingId}", h.MatchingUserHandler)
	authRouter.HandleFunc("/messages", h.MessageHandler).Methods("POST")
	authRouter.HandleFunc("/messages/{messageId}", h.MessageHandler)
	authRouter.HandleFunc("/question_cards", h.QuestionCardHandler)

	return middleware.CORS(router)
}
