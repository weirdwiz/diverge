package data

import (
	"context"
	"errors"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	uuid "github.com/satori/go.uuid"
)

// User : struct for the user
type User struct {
	Username  string    `json:"username" bson:"username"`
	Name      string    `json:"name" bson:"name"`
	ID        string    `json:"userID" bson:"userID"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"createOn" bson:"createdAt"`
	Password  string
}

// Session struct for the User sessions
type Session struct {
	UUID      string    `bson:"uuid"`
	userID    string    `bson:"userID"`
	CreatedAt time.Time `bson:"createdAt"`
}

//CreateSession creates a new user session
func (u *User) CreateSession() (Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s := Session{
		UUID:      uuid.NewV4().String(),
		userID:    u.ID,
		CreatedAt: time.Now(),
	}
	_, err := sessions.InsertOne(ctx, s)
	if err != nil {
		return Session{}, err
	}
	return s, nil
}

// Check : Checks the user session
func (u *User) Check() (valid bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var s Session
	err = sessions.FindOne(ctx, bson.M{
		"userID": u.ID,
	}).Decode(&s)
	if err != nil {
		valid = false
		return
	} else if s.userID == u.ID {
		valid = true
	}
	return
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
