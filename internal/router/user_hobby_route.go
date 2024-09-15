package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewUserHobbyRouter(db *sql.DB, router *mux.Router) {
	uhr := repositories.NewUserHobbyRepo(db)
	uhu := usecases.NewUserHobbyUsecase(uhr)
	uhc := controllers.NewUserHobbyController(uhu)

	router.HandleFunc("/user-hobbies", uhc.CreateUserHobby).Methods("POST")
	router.HandleFunc("/user-hobbies", uhc.UpdateUserHobby).Methods("PUT")
	router.HandleFunc("/user-hobbies/{userId}", uhc.GetAllUserHobby).Methods("GET")
	router.HandleFunc("/user-hobbies/{userId}/category-percentages", uhc.GetUserCategoryPercentages).Methods("GET")
}
