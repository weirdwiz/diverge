package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/weirdwiz/diverge/data"
	"golang.org/x/crypto/bcrypt"
)

// ShowRegisterForm [/register] handles the view of the registeration page
func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	errMsg := r.URL.Query()["err"]
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.ExecuteTemplate(w, "register.html", errMsg)
}

// SubmitRegisterForm [/register/post] handler the submition of the registration page
func SubmitRegisterForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	if r.Form["password"][0] == r.Form["retype-password"][0] {
		bPass := []byte(r.Form["password"][0])
		p, err := bcrypt.GenerateFromPassword(bPass, 1)
		u := &data.User{
			Username:  r.Form["username"][0],
			Name:      r.Form["name"][0],
			Email:     r.Form["email"][0],
			CreatedAt: time.Now(),
			ID:        uuid.NewV4().String(),
			Password:  string(p),
		}

		err = u.Create()
		if err != nil {
			// Handle error
			danger(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/register?err=Passwordnotthesame", http.StatusForbidden)
	}

}
