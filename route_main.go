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
	_, err := session(writer, request)
	launchtime, _ := time.Parse(time.RFC822, "04 Feb 19 15:30 UTC")
	if err != nil {
		http.Redirect(writer, request, "/login", http.StatusFound)
	} else {
		if time.Now().Before(launchtime) {
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
