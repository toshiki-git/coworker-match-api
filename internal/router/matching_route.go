package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewMatchingRouter(db *sql.DB, router *mux.Router) {
	mr := repositories.NewMatchingRepo(db)
	ur := repositories.NewUserRepo(db)
	mu := usecases.NewMatchingUsecase(mr, ur)
	mc := controllers.NewMatchingController(mu)

	router.HandleFunc("/matchings", mc.GetMatchings).Methods("GET")
	router.HandleFunc("/matchings/{matchingId}", mc.GetMatchingUser).Methods("GET")

}
