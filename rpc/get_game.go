package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type GetGameRequest struct {
	GameID string `json:"game_id"`
}

func (r *RPC) GetGame(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	gameRequest := &GetGameRequest{}
	err := r.bindJSON(req, gameRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if gameRequest.GameID == "" {
		r.handleError(w, req, crossesErr.New("missing_game_id", nil))
		return
	}

	game, err := r.app.GetGame(ctx, gameRequest.GameID)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, game)
}
