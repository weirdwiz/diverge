package main

import (
	"fmt"
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
	r.HandleFunc("/login", ShowLogin)
	r.HandleFunc("/authenticate", Authenticate).Methods("POST")
	return r
}

func main() {
	r := newRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO")
}
