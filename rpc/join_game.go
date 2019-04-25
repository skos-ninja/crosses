package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type JoinGameRequest struct {
	GameID string `json:"game_id"`
	UserID string `json:"user_id"`
}

func (r *RPC) JoinGame(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	joinRequest := &JoinGameRequest{}
	err := r.bindJSON(req, joinRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if joinRequest.GameID == "" {
		r.handleError(w, req, crossesErr.New("missing_game_id", nil))
		return
	}

	if joinRequest.UserID == "" {
		r.handleError(w, req, crossesErr.New("missing_user_id", nil))
		return
	}

	game, err := r.app.JoinGame(ctx, joinRequest.GameID, joinRequest.UserID)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, game)
}
