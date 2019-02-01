package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func showQuestion(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		u, err := s.User()
		if err != nil {
			danger(err)
			errorMessage(w, r, "There was a problem")
		}
		q, err := u.GetQuestion()
		if err != nil {
			danger(err)
			errorMessage(w, r, "There was a 2")
		}
		generateHTML(w, template.HTML(q), "question", "private.navbar", "layout", "footer")
	}
}

func checkQuestion(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		u, err := s.User()
		if err != nil {
			danger(err)
			errorMessage(w, r, "There was a problem")
		}
		u.LogAnswer(r.PostFormValue("answer"))
		expected, err := u.GetAnswer()
		if err != nil {
			danger(err)
			errorMessage(w, r, "There was a problem")
		}
		if expected != r.PostFormValue("answer") {
			fmt.Println(expected, r.PostFormValue("answer"), "wrong andwer")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println(expected, r.PostFormValue("answer"), "correect anbswer")

			u.NextLevel()
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
}

// func showQuestion(w http.ResponseWriter, r *http.Request) {
// 	generateHTML(w, nil, "private.navbar", "layout", "index", "footer")
// }
