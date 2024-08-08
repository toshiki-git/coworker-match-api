package router

import (
	"database/sql"

	"github.com/coworker-match-api/internal/controllers"
	"github.com/coworker-match-api/internal/repositories"
	"github.com/coworker-match-api/internal/usecases"
	"github.com/gorilla/mux"
)

func NewMessageRouter(db *sql.DB, router *mux.Router) {
	mr := repositories.NewMessageRepo(db)
	mu := usecases.NewMessageUsecase(mr)
	mc := controllers.NewMessageController(mu)

	router.HandleFunc("/messages/{matchingId}", mc.CreateMessage).Methods("POST")
	router.HandleFunc("/messages/{matchingId}", mc.GetMessages).Methods("GET")
	router.HandleFunc("/messages/{messageId}", mc.UpdateMessage).Methods("PUT")
}
