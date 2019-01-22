package main

import "net/http"

// func showQuestion(w http.ResponseWriter, r *http.Request) {
// 	s, err := session(w, r)
// 	if err != nil {
// 		http.Redirect(w, r, "/login", http.StatusFound)
// 	} else {
// 		u, err := s.User()
// 		if err != nil {
// 			danger(err)
// 			errorMessage(w, r, "There was a problem")
// 		}
// 		q, err := u.GetQuestion()
// 		if err != nil {
// 			danger(err)
// 			errorMessage(w, r, "There was a problem")
// 		}
// 		generateHTML(w, q, "index", "private.navbar", "layout")
// 	}
// }

func showQuestion(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "private.navbar", "layout", "index")
}
