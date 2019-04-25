package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type GetPlayerRequest struct {
	UserID string `json:"user_id"`
}

func (r *RPC) GetPlayer(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	playerRequest := &GetPlayerRequest{}
	err := r.bindJSON(req, playerRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if playerRequest.UserID == "" {
		r.handleError(w, req, crossesErr.New("missing_user_id", nil))
		return
	}

	player, err := r.app.GetPlayer(ctx, playerRequest.UserID)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, player)
}
