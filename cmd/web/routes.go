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

	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireRoleTeacher).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireRoleTeacher).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))

	mux.Get("/toys/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createToyForm))
	mux.Post("/toys/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createToy))

	mux.Get("/toys", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showToys))
	mux.Get("/toys/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showToy))

	mux.Get("/department/create", dynamicMiddleware.Append(app.requireRoleTeacher).ThenFunc(app.createDepartmentForm))
	mux.Post("/department/create", dynamicMiddleware.Append(app.requireRoleTeacher).ThenFunc(app.createDepartment))

	mux.Get("/students", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.students))
	mux.Get("/department", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.Departments))
	mux.Get("/department/:id", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.showDepartment))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// Routes for admin only
	mux.Get("/admin", dynamicMiddleware.Append(app.requireRoleAdmin).ThenFunc(app.adminPanelForm))
	mux.Post("/admin", dynamicMiddleware.Append(app.requireRoleAdmin).ThenFunc(app.adminPanel))

	// Leave the static files route unchanged.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)
}
