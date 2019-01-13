package main

import (
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
	http.ListenAndServe(":8000",
		csrf.Protect(
			[]byte("32-byte-long-auth-key"),
			csrf.Secure(false), // Pass it *to* this constructor
		)(r))
}
