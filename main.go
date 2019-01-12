package main

import (
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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
	key = "32-bit-key"
	CSRF := csrf.Protect([]byte(key))
	log.Fatal(http.ListenAndServe(":8080", CSRF(r)))
}
