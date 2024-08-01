package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippetsList, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/nav.tmpl",
	}
	data := app.newTemplateData(r)
	data.Snippets = snippetsList

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal server error", 500)
	}
	er := ts.ExecuteTemplate(w, "base", data)
	if er != nil {
		app.serverError(w, er)
		http.Error(w, "Internal server error", 500)
	}
}

func (app *application) view(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.notfound(w)
		return
	}

	Snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notfound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	flash := app.sessionManager.PopString(r.Context(), "flash")
	data := app.newTemplateData(r)
	data.Snippet = Snippet

	data.Flash = flash
	// plus the base layout and navigation partial that we made earlier.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	// Parse the template files...
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// And then execute them. Notice how we are passing in the snippet
	// data (a models.Snippet struct) as the final parameter?
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) create(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/nav.tmpl",
		"./ui/html/pages/create.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}
	er := ts.ExecuteTemplate(w, "base", nil)
	if er != nil {
		app.serverError(w, er)
	}

}

type snippetCreateForm struct {
	Title     string `form:"title"`
	Content   string `form:"content"`
	Expires   int    `form:"expires"`
	Validator `form:"-"`
}

func (app *application) createpost(w http.ResponseWriter, r *http.Request) {
	var form snippetCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(form.isBlank(form.Title), "title", "The title cannot be blank")
	form.CheckField(form.maxChars(100, form.Title), "title", "The title must be under 100 characters")
	form.CheckField(form.isBlank(form.Content), "content", "The content cannot be blank")
	form.CheckField(form.permitted(form.Expires, 1, 7, 365), "expires", "The expiration time must be 1, 7 or 365 days")

	if form.valid() == false {
		for key, message := range form.fromErrors {
			w.Write([]byte(key + ": " + message + "\n"))
		}
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}

type userSignupForm struct {
	Name      string `form:"name"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	Validator `form:"-"`
}

func (app *application) signupGet(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/nav.tmpl",
		"./ui/html/pages/signup.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}
	er := ts.ExecuteTemplate(w, "base", nil)
	if er != nil {
		app.serverError(w, er)
	}

}
func (app *application) signupPost(w http.ResponseWriter, r *http.Request) {
	var form userSignupForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(form.isBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(form.isBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(form.matches(EmailRX, form.Email), "email", "This field must be a valid email address")
	form.CheckField(form.isBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(form.minChars(8, form.Password), "password", "This field must be at least 8 characters long")

	if !form.valid() {
		data := app.newTemplateData(r)
		data.Form = form
		// for key, message := range form.fromErrors {
		// 	w.Write([]byte(key + ": " + message + "\n"))
		// }
		// return

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/nav.tmpl",
			"./ui/html/pages/signup.tmpl",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
		}
		er := ts.ExecuteTemplate(w, "base", data)
		if er != nil {
			app.serverError(w, er)
		}
	}
	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("INSERT: Name: " + form.Name + " | Email: " + form.Email + " | Password: " + form.Password)
	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

type userLoginForm struct {
	name      string `form:"name"`
	password  string `form:"password"`
	Validator `form:"-"`
}

func (app *application) loginGet(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userLoginForm{}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/nav.tmpl",
		"./ui/html/pages/login.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}
	er := ts.ExecuteTemplate(w, "base", data)
	if er != nil {
		app.serverError(w, er)
	}
}

func (app *application) loginPost(w http.ResponseWriter, r *http.Request) {
	var form userLoginForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
func (app *application) logoutPost(w http.ResponseWriter, r *http.Request) {

}
