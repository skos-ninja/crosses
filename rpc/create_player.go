package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type CreatePlayerRequest struct {
	Name string `json:"name"`
}

func (r *RPC) CreatePlayer(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	playerRequest := &CreatePlayerRequest{}
	err := r.bindJSON(req, playerRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if playerRequest.Name == "" {
		r.handleError(w, req, crossesErr.New("missing_user_name", nil))
		return
	}

	player, err := r.app.CreatePlayer(ctx, playerRequest.Name)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, player)
}
