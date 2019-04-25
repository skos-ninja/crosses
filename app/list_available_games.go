package app

import (
	"context"

	"crosses/services/game/models"
)

// ListAvailableGames will return any game that has only 1 player and is free to join
func (app *App) ListAvailableGames(ctx context.Context) ([]*models.Game, error) {
	games, err := app.gameService.ListAvailableGames(ctx)
	if err != nil {
		return nil, err
	}

	return games, nil
}
