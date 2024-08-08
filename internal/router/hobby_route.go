package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewHobbyRouter(db *sql.DB, router *mux.Router) {
	hr := repositories.NewHobbyRepo(db)
	hu := usecases.NewHobbyUsecase(hr)
	hc := controllers.NewHobbyController(hu)

	router.HandleFunc("/hobbies", hc.GetAllHobby).Methods("GET")
}
