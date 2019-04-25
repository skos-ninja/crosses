package app

import (
	"context"

	crossesErr "crosses/err"
	"crosses/services/board"
	"crosses/services/game/models"
)

func (app *App) GetGameState(ctx context.Context, gameID string) ([]*board.Position, error) {
	game, err := app.gameService.GetGame(ctx, gameID)
	if err != nil {
		return nil, err
	}

	brd, err := app.getBoardForGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return brd.GetState(), nil
}

func (app *App) getBoardForGame(ctx context.Context, game *models.Game) (*board.Board, error) {
	turns, err := app.gameService.GetTurns(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	positions := make([]*board.Position, 0)
	for _, turn := range turns {
		position, err := convertTurnToPosition(game, turn)
		if err != nil {
			return nil, err
		}

		positions = append(positions, position)
	}

	brd := board.NewBoard()
	err = brd.SetState(positions)
	if err != nil {
		return nil, err
	}

	return &brd, nil
}

func convertTurnToPosition(game *models.Game, turn *models.Turn) (*board.Position, error) {
	state := board.Empty
	if turn.PlayerID == game.Player1ID {
		state = board.X
	} else if turn.PlayerID == game.Player2ID {
		state = board.O
	} else {
		return nil, crossesErr.New("invalid_player_for_turn", nil)
	}

	return &board.Position{
		X:     turn.XAxis,
		Y:     turn.YAxis,
		State: state,
	}, nil
}
