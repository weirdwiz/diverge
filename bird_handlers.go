package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Bird is the structure for the Birds
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	birds, err := store.GetBirds()

	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	err = store.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/static/", http.StatusFound)
}
