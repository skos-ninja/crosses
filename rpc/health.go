package rpc

import (
	"net/http"
)

func (r *RPC) HealthCheck(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	healthy := r.app.CheckHealth(ctx)
	if !healthy {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
	return
}
