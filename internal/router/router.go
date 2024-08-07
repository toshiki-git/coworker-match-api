package router

import (
	"database/sql"
	"net/http"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/middleware"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func InitRouter(db *sql.DB) http.Handler {
	userRepo := repositories.NewUserRepo(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)

	hobbyRepo := repositories.NewHobbyRepo(db)
	hobbyUsecase := usecases.NewHobbyUsecase(hobbyRepo)
	hobbyController := controllers.NewHobbyController(hobbyUsecase)

	userHobbyRepo := repositories.NewUserHobbyRepo(db)
	userHobbyUsecase := usecases.NewUserHobbyUsecase(userHobbyRepo)
	userHobbyController := controllers.NewUserHobbyController(userHobbyUsecase)

	questionCardRepo := repositories.NewQuestionCardRepo(db)
	questionCardUsecase := usecases.NewQuestionCardUsecase(questionCardRepo)
	questionCardController := controllers.NewQuestionCardController(questionCardUsecase)

	matchingQuestionRepo := repositories.NewMatchingQuestionRepo(db)
	matchingQuestionUsecase := usecases.NewMatchingQuestionUsecase(matchingQuestionRepo)
	matchingQuestionController := controllers.NewMatchingQuestionController(matchingQuestionUsecase)

	matchingRepo := repositories.NewMatchingRepo(db)
	matchingUsecase := usecases.NewMatchingUsecase(matchingRepo)
	matchingController := controllers.NewMatchingController(matchingUsecase)

	messageRepo := repositories.NewMessageRepo(db)
	messageUsecase := usecases.NewMessageUsecase(messageRepo)
	messageController := controllers.NewMessageController(messageUsecase)

	router := mux.NewRouter()

	// 非認証エンドポイント
	router.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	// 認証エンドポイントのグループ化
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.Auth)

	authRouter.HandleFunc("/users/exists", userController.IsUserExist).Methods("GET")
	authRouter.HandleFunc("/users", userController.CreateUser).Methods("POST")
	authRouter.HandleFunc("/users/{userId}", userController.GetUserById).Methods("GET")
	authRouter.HandleFunc("/users/{userId}", userController.UpdateUser).Methods("PUT")

	authRouter.HandleFunc("/hobbies", hobbyController.GetAllHobby).Methods("GET")

	authRouter.HandleFunc("/user-hobbies", userHobbyController.CreateUserHobby).Methods("POST")
	authRouter.HandleFunc("/user-hobbies", userHobbyController.UpdateUserHobby).Methods("PUT")
	authRouter.HandleFunc("/user-hobbies/{userId}", userHobbyController.GetAllUserHobby).Methods("GET")

	authRouter.HandleFunc("/matching-questions", matchingQuestionController.CreateMatching).Methods("POST")
	authRouter.HandleFunc("/matching-questions", matchingQuestionController.GetMatchingQuestion).Methods("GET")

	authRouter.HandleFunc("/question-cards/{matchingId}", questionCardController.GetQuestionCards).Methods("GET")

	authRouter.HandleFunc("/matchings", matchingController.GetMatchings).Methods("GET")
	authRouter.HandleFunc("/matchings/{matchingId}", matchingController.GetMatchingUser).Methods("GET")

	authRouter.HandleFunc("/messages/{matchingId}", messageController.CreateMessage).Methods("POST")
	authRouter.HandleFunc("/messages/{matchingId}", messageController.GetMessages).Methods("GET")
	authRouter.HandleFunc("/messages/{messageId}", messageController.UpdateMessage).Methods("PUT")

	return middleware.CORS(router)
}
