package blacksmith

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//routes create a webserver using chi
func (bls *Blacksmith) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	//only in debug mode
	if bls.Debug {
		mux.Use(middleware.Logger)
	}
	mux.Use(middleware.Recoverer)
	//blacksmith middleware to load session configured parameters
	mux.Use(bls.SessionLoad)

	return mux

}
