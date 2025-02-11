package services

import (
	"game-service/internal/db/repositories"
)

type GameServ struct {
	gameRepo repositories.GameRepository
}

func NewGameService(gameRepo repositories.GameRepository) GameServices {
	return &GameServ{gameRepo}
}
