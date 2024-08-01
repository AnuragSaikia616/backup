package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter" // New import
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.view))
	router.Handler(http.MethodGet, "/snippet/create", dynamic.ThenFunc(app.create))
	router.Handler(http.MethodPost, "/snippet/create", dynamic.ThenFunc(app.createpost))

	//authentication routes
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.signupGet))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.signupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.loginGet))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.loginPost))
	router.Handler(http.MethodPost, "/user/logout", dynamic.ThenFunc(app.logoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
