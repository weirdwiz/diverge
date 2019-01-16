package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/weirdwiz/diverge/data"
)

// ShowLogin [/login] displays the login page
func ShowLogin(w http.ResponseWriter, r *http.Request) {
	errMsg := r.URL.Query()["err"]
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.ExecuteTemplate(w, "login.html", errMsg)
}

// Authenticate [/authenticate] accepts the post request for login
func Authenticate(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	user, err := data.GetUserByUsername(r.PostFormValue("username"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.PostFormValue("password"))); err == nil {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}
