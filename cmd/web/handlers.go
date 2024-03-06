package main

import (
	"errors"
	"fmt"
	"github.com/oynaToys/pkg/forms"
	"github.com/oynaToys/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	s, err := app.toys.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "toys.page.tmpl", &templateData{
		Toys: s,
	})
}

func (app *application) createFeedbackForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.feedback.page.tmpl", &templateData{
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
}

func (app *application) createFeedback(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "stars")

	if !form.Valid() {
		app.render(w, r, "create.feedback.page.tmpl", &templateData{Form: form})
		return
	}
	id, err := app.feedbacks.InsertFeedback(form.Get("name"), form.Get("content"), form.Get("stars"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Feedback successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/feedback/%d", id), http.StatusSeeOther)
}

func (app *application) Feedbacks(w http.ResponseWriter, r *http.Request) {
	s, err := app.feedbacks.ShowFeedbacks()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "feedbacks.page.tmpl", &templateData{
		Feedbacks: s,
	})
}

func (app *application) showFeedback(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.feedbacks.GetFeedback(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "show.feedback.page.tmpl", &templateData{
		Feedback: s,
	})
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	// Parse the form data.
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Validate the form contents using the form helper we made earlier.
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	// If there are any errors, redisplay the signup form.
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Check whether the credentials are valid. If they're not, add a generic error
	// message to the form failures map and re-display the login page.
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Add the ID of the current user to the session, so that they are now 'logged
	// in'.
	app.session.Put(r, "authenticatedUserID", id)
	// Redirect the user to the create snippet page.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	app.session.Remove(r, "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) adminPanelForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "admin.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) adminPanel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("email")
	form.MatchesPattern("email", forms.EmailRX)

	if !form.Valid() {
		app.render(w, r, "admin.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.users.DeleteUser(form.Get("email"))

	if err != nil {
		if errors.Is(err, models.ErrEmailDoesNotExist) {
			form.Errors.Add("email", "No user with such email")
			app.render(w, r, "admin.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.session.Put(r, "flash", "User was deleted successfully!")

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (app *application) showToys(w http.ResponseWriter, r *http.Request) {
	t, err := app.toys.GetToys()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "toys.page.tmpl", &templateData{
		Toys: t,
	})

}

func (app *application) showToy(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.toys.GetToy(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "toy.page.tmpl", &templateData{
		Toy: s,
	})

}

func (app *application) createToy(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "description", "tokens")
	form.MaxLength("title", 100)
	if !form.Valid() {
		app.render(w, r, "create.toy.page.tmpl", &templateData{Form: form})
		return
	}
	id, err := app.toys.InsertToy(form.Get("name"), form.Get("description"), form.Get("tokens"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Toy successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/toys/%d", id), http.StatusSeeOther)
}

func (app *application) createToyForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.toy.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}
