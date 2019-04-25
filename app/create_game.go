package app

import (
	"context"

	"crosses/services/game/models"
)

// CreateGame will create a new game and attach the user ID passed in as player 1
func (app *App) CreateGame(ctx context.Context, userID string) (*models.Game, error) {
	// Ensure user exists before we create
	user, err := app.userService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	game, err := app.gameService.CreateGame(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return game, nil
}
