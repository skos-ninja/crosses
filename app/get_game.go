package app

import (
	"context"

	"crosses/services/game/models"
)

// GetGame will return basic details of the game based off of ID.
func (app *App) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	game, err := app.gameService.GetGame(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return game, nil
}
