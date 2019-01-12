package main

import "errors"

// User : struct for the user
type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	ID       string `json:"userID"` // UUID TODO
	Email    string `json:"email"`
	CreateOn string `json:"createdon"` // TIME TODO
}

func (u *User) register() error {
	//todo add to db, maybe mongo db
	return errors.New("TODO")
}
