package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var key string

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/register", ShowRegisterForm)
	r.HandleFunc("/register/post", SubmitRegisterForm).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	log.Fatal(http.ListenAndServe(":6969", r))
}
