package rpc

import (
	"net/http"
)

func (r *RPC) ListAvailableGames(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	games, err := r.app.ListAvailableGames(ctx)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, games)
}
