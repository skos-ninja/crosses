package app

import (
	"context"

	"crosses/services/user/models"
)

// CreatePlayer will create a new player with the given name.
// This player ID is used in requests to identify the user in the game api
func (app *App) CreatePlayer(ctx context.Context, name string) (*models.Player, error) {
	player, err := app.userService.CreateUser(ctx, name)
	if err != nil {
		return nil, err
	}

	return player, nil
}
