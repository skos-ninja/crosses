package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type GetGameStateRequest struct {
	GameID string `json:"game_id"`
}

func (r *RPC) GetGameState(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	gameRequest := &GetGameStateRequest{}
	err := r.bindJSON(req, gameRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if gameRequest.GameID == "" {
		r.handleError(w, req, crossesErr.New("missing_game_id", nil))
		return
	}

	turns, err := r.app.GetGameState(ctx, gameRequest.GameID)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, turns)
}
