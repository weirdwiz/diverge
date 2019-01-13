package data

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

// User : struct for the user
type User struct {
	Username string    `json:"username" bson:"username"`
	Name     string    `json:"name" bson:"name"`
	ID       uuid.UUID `json:"userID" bson:"userID"` // UUID TODO
	Email    string    `json:"email" bson:"email"`
	CreateOn time.Time `json:"createOn" bson:"createOn"` // TODO TIME CHOOSE A BETTER NAME
}

// Create a new user
func (u *User) Create() error {

	collection := client.Database("diverge").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	return nil
}
