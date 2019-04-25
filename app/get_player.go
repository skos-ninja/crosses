package app

import (
	"context"

	"crosses/services/user/models"
)

// GetPlayer will return basic details of the player based off of ID
func (app *App) GetPlayer(ctx context.Context, userID string) (*models.Player, error) {
	player, err := app.userService.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return player, nil
}
