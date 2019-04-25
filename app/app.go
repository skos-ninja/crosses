package app

import (
	"crosses/db"

	"crosses/services/game"
	"crosses/services/user"
)

// App contains all the code to actually handle a request
type App struct {
	db          *db.DB
	gameService game.Service
	userService user.Service
}

// New creates a new instance of the App
func New(db *db.DB, gameService game.Service, userService user.Service) *App {
	return &App{
		db:          db,
		gameService: gameService,
		userService: userService,
	}
}
