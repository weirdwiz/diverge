package main

import (
	"net/http"
	"time"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error", "footer")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error", "footer")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	s, err := session(writer, request)
	launchtime, _ := time.Parse(time.RFC822, "04 Feb 19 11:50 UTC")
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
	} else {
		user, err := s.User()
		if err != nil {
			danger("cannot detect username")
		}
		if time.Now().Before(launchtime) && user.Username != "weirdwiz" {
			generateHTML(writer, nil, "index", "layout", "private.navbar", "footer")
		} else {
			http.Redirect(writer, request, "/play", http.StatusFound)
		}
	}
}

func showRules(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		errorMessage(w, r, "not logged in")
	}
	generateHTML(w, nil, "layout", "private.navbar", "rules", "footer")
}
