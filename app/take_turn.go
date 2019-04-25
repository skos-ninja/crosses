package app

import (
	"context"

	crossesErr "crosses/err"
	"crosses/services/board"
	"crosses/services/game/models"
)

type TakeTurnResponse struct {
	Game *models.Game `json:"game"`
	Turn *models.Turn `json:"turn"`
}

// TakeTurn will validate a turn and then commit to DB and then perform a win check
func (app *App) TakeTurn(ctx context.Context, gameID string, userID string, x, y int) (*TakeTurnResponse, error) {
	game, err := app.gameService.GetGame(ctx, gameID)
	if err != nil {
		return nil, err
	}

	brd, err := app.getBoardForGame(ctx, game)
	if err != nil {
		return nil, err
	}

	player := board.Empty
	if game.Player1ID == userID {
		player = board.X
	} else if game.Player2ID == userID {
		player = board.O
	} else {
		return nil, crossesErr.New("player_not_in_game", nil)
	}

	if brd.CalculateCurrentPlayersTurn() != player {
		return nil, crossesErr.New("player_not_current_turn", nil)
	}

	if !brd.CanTakeTurn(x, y) {
		return nil, crossesErr.New("position_already_taken", nil)
	}

	pos := board.Position{
		X:     x,
		Y:     y,
		State: player,
	}

	err = brd.SetPosition(&pos)
	if err != nil {
		return nil, err
	}

	turn, err := app.gameService.CreateTurn(ctx, game.ID, userID, x, y)
	if err != nil {
		return nil, err
	}

	winner := brd.CalculateWinnerFromPosition(x, y)
	if winner != board.Empty {
		winnerID := ""
		if winner == board.X {
			winnerID = game.Player1ID
		} else {
			winnerID = game.Player2ID
		}

		game, err = app.gameService.DeclareWinner(ctx, game.ID, winnerID)
	}

	return &TakeTurnResponse{
		Game: game,
		Turn: turn,
	}, nil
}
