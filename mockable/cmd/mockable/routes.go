package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *config) initRoutes() http.Handler {
	mux := chi.NewRouter()

	return mux

}
