package main

import (
	"net/http"
	"time"
)

func main() {
	p("Labyrinth", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_play.go
	mux.HandleFunc("/play", showQuestion)
	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/leaderboard", showLeaderBoard)
	mux.HandleFunc("/rules", showRules)
	mux.HandleFunc("/checkAnswer", checkQuestion)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
