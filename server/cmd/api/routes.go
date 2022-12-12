package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/music/:id", app.getOneMusic)
	router.HandlerFunc(http.MethodGet, "/music", app.getAllMusic)

	router.HandlerFunc(http.MethodPost, "/music/create", app.AddMusic)
	router.HandlerFunc(http.MethodPost, "/music/edit/", app.EditMusic)
	router.HandlerFunc(http.MethodPost, "/music/delete/", app.DeleteMusic)

	return app.enableCORS(router)
}
