package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
	Name   string
}

var MONGO_URI = "mongodb://localhost:27017"

func ConnectMongo() (*DB, error) {
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Printf("Captured command: %v\n", evt.Command)
		},
	}

	// if uri := os.Getenv("MONGO_URI"); uri != "" {
	// 	MONGO_URI = uri
	// }

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI).SetMonitor(monitor))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	return &DB{Client: client, Name: "testdb"}, nil
}
