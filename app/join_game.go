package app

import (
	"context"

	"crosses/services/game/models"
)

// JoinGame will join the passed user ID to the given game
func (app *App) JoinGame(ctx context.Context, gameID string, userID string) (*models.Game, error) {
	user, err := app.userService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	game, err := app.gameService.JoinGame(ctx, gameID, user.ID)
	if err != nil {
		return nil, err
	}

	return game, nil
}
