package data

import (
	"context"
	"errors"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

// User : struct for the user
type User struct {
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	ID       string    `json:"userID" bson:"userID"`
	Email    string    `json:"email" bson:"email"`
	CreateOn time.Time `json:"createOn" bson:"createOn"`
}

// Create a new user
func (u *User) Create() error {
	// Verify that the username is unique the repeated password is same

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := users.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByUsername returns a pointer to the user object by searching by username
func GetUserByUsername(username string) (*User, error) {
	var u *User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := users.FindOne(ctx, bson.M{
		"username": u.Username,
	}).Decode(u)
	if err != nil {
		return nil, err
	}
	return u, errors.New("username not fucking found")
}
