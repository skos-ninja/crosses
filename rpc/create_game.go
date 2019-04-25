package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type CreateGameRequest struct {
	UserID string `json:"user_id"`
}

func (r *RPC) CreateGame(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	gameRequest := &CreateGameRequest{}
	err := r.bindJSON(req, gameRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if gameRequest.UserID == "" {
		r.handleError(w, req, crossesErr.New("missing_user_id", nil))
		return
	}

	game, err := r.app.CreateGame(ctx, gameRequest.UserID)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, game)
}
