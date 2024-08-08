package controllers

import (
	"net/http"

	"github.com/coworker-match-api/internal/usecases"
)

type MatchingController struct {
	mu usecases.IMatchingUsecase
}

func NewMatchingController(mu usecases.IMatchingUsecase) *MatchingController {
	return &MatchingController{mu: mu}
}

func (mc *MatchingController) GetMatchings(w http.ResponseWriter, r *http.Request) {
}

func (mc *MatchingController) GetMatchingUser(w http.ResponseWriter, r *http.Request) {
}
