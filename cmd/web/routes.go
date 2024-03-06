package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// Create a new middleware chain containing the middleware specific to
	// our dynamic application routes. For now, this chain will only contain
	// the session middleware but we'll add more to it later.
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)
	mux := pat.New()
	// Update these routes to use the new dynamic middleware chain followed
	// by the appropriate handler function.

	// Routes
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	
	mux.Get("/toys/create", dynamicMiddleware.Append(app.requireRoleAdmin).ThenFunc(app.createToyForm))
	mux.Post("/toys/create", dynamicMiddleware.Append(app.requireRoleAdmin).ThenFunc(app.createToy))

	mux.Get("/toys", dynamicMiddleware.ThenFunc(app.showToys))
	mux.Get("/toys/:id", dynamicMiddleware.ThenFunc(app.showToy))

	mux.Get("/feedback/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createFeedbackForm))
	mux.Post("/feedback/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createFeedback))

	mux.Get("/feedbacks", dynamicMiddleware.ThenFunc(app.Feedbacks))
	mux.Get("/feedback/:id", dynamicMiddleware.ThenFunc(app.showFeedback))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// Routes for admin only
	mux.Get("/admin", dynamicMiddleware.Append(app.requireRoleAdmin).ThenFunc(app.adminPanelForm))
	mux.Post("/admin", dynamicMiddleware.Append(app.requireRoleAdmin).ThenFunc(app.adminPanel))

	// Leave the static files route unchanged.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
