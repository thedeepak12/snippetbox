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

	router.Get("/", app.home)
	router.Get("/snippet/view/{id}", app.snippetView)
	router.Get("/snippet/create", app.snippetCreate)
	router.Post("/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
