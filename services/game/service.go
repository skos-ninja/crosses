package game

import (
	"context"
	"database/sql"
	"time"

	"crosses/db"
	crossErr "crosses/err"
	"crosses/services/game/models"

	"github.com/cuvva/ksuid"
)

// Service handles managing games and the taking of turns
type Service interface {
	CreateGame(ctx context.Context, userID string) (*models.Game, error)
	GetGame(ctx context.Context, gameID string) (*models.Game, error)
	ListAvailableGames(ctx context.Context) ([]*models.Game, error)
	JoinGame(ctx context.Context, gameID string, userID string) (*models.Game, error)

	GetTurns(ctx context.Context, gameID string) ([]*models.Turn, error)
	CreateTurn(ctx context.Context, gameID string, userID string, xAxis, yAxis int) (*models.Turn, error)
	DeclareWinner(ctx context.Context, gameID, winnerID string) (*models.Game, error)
}

type gameService struct {
	db *db.DB
}

// NewGameService creates a new GameService instance
func NewGameService(db *db.DB) Service {
	return &gameService{db: db}
}

func (svc *gameService) GetGame(ctx context.Context, gameID string) (*models.Game, error) {
	row := svc.db.QueryRowWithCtx(ctx, "SELECT id, player1_id, player2_id, started_at, finished_at, winning_player FROM games WHERE id=$1", gameID)
	if row != nil {
		id := ""
		var player1 sql.NullString
		var player2 sql.NullString
		var startedAt *time.Time
		var finishedAt *time.Time
		var winningPlayer sql.NullString

		err := row.Scan(&id, &player1, &player2, &startedAt, &finishedAt, &winningPlayer)
		if err != nil {
			return nil, err
		}

		return &models.Game{
			ID:            id,
			Player1ID:     player1.String,
			Player2ID:     player2.String,
			StartedAt:     startedAt,
			FinishedAt:    finishedAt,
			WinningPlayer: winningPlayer.String,
		}, nil
	}

	return nil, crossErr.New("game_not_found", nil)
}

func (svc *gameService) ListAvailableGames(ctx context.Context) ([]*models.Game, error) {
	rows, err := svc.db.QueryWithCtx(ctx, "SELECT id, player1_id, player2_id FROM games WHERE player1_id IS NOT NULL AND player2_id IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	games := make([]*models.Game, 0)

	for rows.Next() {
		id := ""
		var player1 sql.NullString
		var player2 sql.NullString

		if err := rows.Scan(&id, &player1, &player2); err != nil {
			return nil, err
		}

		game := models.Game{
			ID:        id,
			Player1ID: player1.String,
			Player2ID: player2.String,
		}

		games = append(games, &game)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (svc *gameService) CreateGame(ctx context.Context, userID string) (*models.Game, error) {
	id := ksuid.Generate("game")

	game := &models.Game{
		ID: id.String(),

		Player1ID: userID,
		Player2ID: "",

		StartedAt:  nil,
		FinishedAt: nil,

		WinningPlayer: "",
	}

	_, err := svc.db.Exec(ctx, "INSERT INTO games (id, player1_id) VALUES ($1, $2)", game.ID, game.Player1ID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (svc *gameService) JoinGame(ctx context.Context, gameID string, userID string) (*models.Game, error) {
	game, err := svc.GetGame(ctx, gameID)
	if err != nil {
		return nil, err
	}

	if game.Player2ID != "" {
		return nil, crossErr.New("game_has_players", nil)
	}

	now := time.Now().UTC()
	game.StartedAt = &now
	game.Player2ID = userID

	_, err = svc.db.Exec(ctx, "UPDATE games SET player2_id=$2, started_at=$3 WHERE id=$1", game.ID, userID, now)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (svc *gameService) GetTurns(ctx context.Context, gameID string) ([]*models.Turn, error) {
	rows, err := svc.db.QueryWithCtx(ctx, "SELECT id, game_id, player_id, x_axis, y_axis, created_at FROM turns WHERE game_id=$1", gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	turns := make([]*models.Turn, 0)

	for rows.Next() {
		turn := models.Turn{}

		if err := rows.Scan(&turn.ID, &turn.GameID, &turn.PlayerID, &turn.XAxis, &turn.YAxis, &turn.CreatedAt); err != nil {
			return nil, err
		}

		turns = append(turns, &turn)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return turns, nil
}

func (svc *gameService) CreateTurn(ctx context.Context, gameID, userID string, xAxis, yAxis int) (*models.Turn, error) {
	id := ksuid.Generate("turn")

	now := time.Now().UTC()
	turn := &models.Turn{
		ID:       id.String(),
		GameID:   gameID,
		PlayerID: userID,

		XAxis: xAxis,
		YAxis: yAxis,

		CreatedAt: &now,
	}

	_, err := svc.db.Exec(ctx, "INSERT INTO turns (id, game_id, player_id, x_axis, y_axis, created_at) VALUES ($1, $2, $3, $4, $5, $6)", turn.ID, turn.GameID, turn.PlayerID, turn.XAxis, turn.YAxis, now)
	if err != nil {
		return nil, err
	}

	return turn, nil
}

func (svc *gameService) DeclareWinner(ctx context.Context, gameID, winnerID string) (*models.Game, error) {
	game, err := svc.GetGame(ctx, gameID)
	if err != nil {
		return nil, err
	}

	if game.WinningPlayer != "" {
		return nil, crossErr.New("game_has_winner", nil)
	}

	now := time.Now().UTC()
	game.FinishedAt = &now
	game.WinningPlayer = winnerID

	_, err = svc.db.Exec(ctx, "UPDATE games SET winning_player=$2, finished_at=$3 WHERE id=$1", game.ID, winnerID, now)
	if err != nil {
		return nil, err
	}

	return game, nil
}
