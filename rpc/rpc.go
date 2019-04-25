package rpc

import (
	"encoding/json"
	"net/http"

	"crosses/app"
	crossesErr "crosses/err"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// RPC is a server for use within mux
type RPC struct {
	logger *logrus.Logger
	router *chi.Mux
	app    *app.App
}

// New creates a new RPC wrapper
func New(logger *logrus.Logger, router *chi.Mux, app *app.App) *RPC {
	return &RPC{
		logger: logger,
		app:    app,
		router: router,
	}
}

// Use is a wrapper for chi's Use func
func (r *RPC) Use(middlewares ...func(http.Handler) http.Handler) {
	r.router.Use(middlewares...)
}

// Get is a wrapper for chi's Get func
func (r *RPC) Get(pattern string, handlerFn http.HandlerFunc) {
	r.router.Get(pattern, handlerFn)
}

// Post is a wrapper for chi's Post func
func (r *RPC) Post(pattern string, handlerFn http.HandlerFunc) {
	r.router.Post(pattern, handlerFn)
}

// Serve starts the HTTP server
func (r *RPC) Serve(port string) {
	r.logger.Info("starting on ", port)

	http.ListenAndServe(port, r.router)
}

func (r *RPC) handleError(w http.ResponseWriter, req *http.Request, err error) {
	log := logrus.NewEntry(r.logger)

	w.Header().Set("Content-Type", "application/json")

	v, ok := err.(crossesErr.E)
	if ok {
		log.Warn(v, v.Meta, v.Reasons)
		w.WriteHeader(500)
	} else {
		log.Warn(err)

		// Hide the original error with an unknown code to stop stacktrace leaks
		v = crossesErr.New(crossesErr.Unknown, nil)

		w.WriteHeader(500)
	}

	json.NewEncoder(w).Encode(v)

	return
}

func (r *RPC) handleJSON(w http.ResponseWriter, output interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(output)

	return
}

func (r *RPC) bindJSON(req *http.Request, output interface{}) error {
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&output)
	if err != nil {
		return crossesErr.New("invalid_body", nil)
	}

	return nil
}
