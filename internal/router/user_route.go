package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewUserRouter(db *sql.DB, router *mux.Router) {
	ur := repositories.NewUserRepo(db)
	uu := usecases.NewUserUsecase(ur)
	uc := controllers.NewUserController(uu)

	router.HandleFunc("/users/exists", uc.IsUserExist).Methods("GET")
	router.HandleFunc("/users", uc.CreateUser).Methods("POST")
	router.HandleFunc("/users/{userId}", uc.GetUserById).Methods("GET")
	router.HandleFunc("/users/{userId}", uc.UpdateUser).Methods("PUT")
}
