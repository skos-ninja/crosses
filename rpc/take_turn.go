package rpc

import (
	"net/http"

	crossesErr "crosses/err"
)

type TakeTurnRequest struct {
	GameID string `json:"game_id"`
	UserID string `json:"user_id"`

	XAxis int `json:"x"`
	YAxis int `json:"y"`
}

func (r *RPC) TakeTurn(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	turnRequest := &TakeTurnRequest{}
	err := r.bindJSON(req, turnRequest)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	if turnRequest.GameID == "" {
		r.handleError(w, req, crossesErr.New("missing_game_id", nil))
		return
	}

	if turnRequest.UserID == "" {
		r.handleError(w, req, crossesErr.New("missing_user_id", nil))
		return
	}

	turn, err := r.app.TakeTurn(ctx, turnRequest.GameID, turnRequest.UserID, turnRequest.XAxis, turnRequest.YAxis)
	if err != nil {
		r.handleError(w, req, err)
		return
	}

	r.handleJSON(w, turn)
}
