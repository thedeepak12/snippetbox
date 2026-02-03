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

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf)
	router.Method("GET", "/", dynamic.ThenFunc(app.home))
	router.Method("GET", "/snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	router.Method("GET", "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Method("POST", "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Method("GET", "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Method("POST", "/user/login", dynamic.ThenFunc(app.userLoginPost))
	
	protected := dynamic.Append(app.requireAuthentication)
	router.Method("GET", "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Method("POST", "/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Method("POST", "/user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
