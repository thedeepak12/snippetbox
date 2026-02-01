package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Method("GET", "/", dynamic.ThenFunc(app.home))
	router.Method("GET", "/snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	router.Method("GET", "/snippet/create", dynamic.ThenFunc(app.snippetCreate))
	router.Method("POST", "/snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
