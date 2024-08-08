package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewQuestionCardRouter(db *sql.DB, router *mux.Router) {
	qcr := repositories.NewQuestionCardRepo(db)
	qcu := usecases.NewQuestionCardUsecase(qcr)
	qcc := controllers.NewQuestionCardController(qcu)

	router.HandleFunc("/question-cards/{matchingId}", qcc.GetQuestionCards).Methods("GET")

}
