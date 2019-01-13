package data

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

func init() {
	// Verify that the username is unique the repeated password is same

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Fucking failed to connect to DB, I am a faggot")
	}
}
