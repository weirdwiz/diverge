package main

import (
	"fmt"
	"net/http"

	"github.com/weirdwiz/labyrinth/data"
)

func showLeaderBoard(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	fmt.Println(err)
	if err != nil {
		errorMessage(w, r, "Not Logged in")
	}
	leaderboard, err := data.GetLeaderBoard()
	if err != nil {
		danger(err)
	}
	generateHTML(w, leaderboard, "layout", "leaderboard", "private.navbar", "footer")
}
