package data

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var users *mongo.Collection

func init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Fucking failed to connect to DB, I am a faggot")
	}
	users = client.Database("diverge").Collection("users")
}
