package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewMatchingQuestionRouter(db *sql.DB, router *mux.Router) {
	mqr := repositories.NewMatchingQuestionRepo(db)
	mqu := usecases.NewMatchingQuestionUsecase(mqr)
	mqc := controllers.NewMatchingQuestionController(mqu)

	router.HandleFunc("/matching-questions", mqc.CreateMatching).Methods("POST")
	router.HandleFunc("/matching-questions", mqc.GetMatchingQuestion).Methods("GET")

}
