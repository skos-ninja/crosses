package main

import (
	"os"
	"time"

	"crosses/app"
	"crosses/db"
	"crosses/rpc"
	"crosses/services/game"
	"crosses/services/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	dbOpts := db.Config{
		ConnectionString: getDbConnString(),
	}

	db, err := db.NewDB(dbOpts)
	if err != nil {
		logger.Fatal(err)
	}

	gameService := game.NewGameService(db)
	userService := user.NewUserService(db)

	router := chi.NewRouter()
	app := app.New(db, gameService, userService)
	r := rpc.New(logger, router, app)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", r.HealthCheck)

	r.Post("/create_player", r.CreatePlayer)
	r.Post("/get_player", r.GetPlayer)

	r.Post("/create_game", r.CreateGame)
	r.Post("/get_game", r.GetGame)
	r.Post("/list_available_games", r.ListAvailableGames)
	r.Post("/join_game", r.JoinGame)
	r.Post("/get_game_state", r.GetGameState)

	r.Post("/take_turn", r.TakeTurn)
	// r.Post("/get_turn", r.GetTurn)

	r.Serve(":8080")
}

func getDbConnString() string {
	v := os.Getenv("PG_OPTS")
	if v == "" {
		return "postgresql://postgres:temppass@192.168.99.100/dev_games?sslmode=disable"
	}

	return v
}
