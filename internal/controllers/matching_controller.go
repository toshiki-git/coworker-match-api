package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type IMatchingController interface {
	GetMatchings(w http.ResponseWriter, r *http.Request)
	GetMatchingUser(w http.ResponseWriter, r *http.Request)
}

type matchingController struct {
	mu usecases.IMatchingUsecase
}

func NewMatchingController(mu usecases.IMatchingUsecase) IMatchingController {
	return &matchingController{mu: mu}
}

func (mc *matchingController) GetMatchings(w http.ResponseWriter, r *http.Request) {
}

func (mc *matchingController) GetMatchingUser(w http.ResponseWriter, r *http.Request) {
}
