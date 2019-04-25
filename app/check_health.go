package app

import (
	"context"
)

// CheckHealth will do a health check of the app's external connections
func (app *App) CheckHealth(ctx context.Context) bool {
	err := app.db.Ping(ctx)
	if err != nil {
		return false
	}

	return true
}
