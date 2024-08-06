package repositories

import (
	models "github.com/coworker-match-api/gen/go"
)

type HobbyRepository interface {
	//CreateHobby()
	GetAllHobby() ([]models.GetHobbyResponseInner, error)
}
