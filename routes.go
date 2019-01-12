package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO")
}

// ShowRegisterForm handles the view of the registeration page
func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.ExecuteTemplate(w, "register.html", map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

// SubmitRegisterForm handler the submition of the registration page
func SubmitRegisterForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	u := &User{
		Username: r.PostForm.Get("username"),
		Name:     r.PostForm.Get("name"),
		Email:    r.PostForm.Get("email"),
	}

	err = u.register()
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
}
