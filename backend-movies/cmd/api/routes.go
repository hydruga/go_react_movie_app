package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	return app.enableCORS(router)
}
