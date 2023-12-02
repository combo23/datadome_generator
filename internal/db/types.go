package db

import (
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	MongoURI string
	client   *mongo.Client
	mu       sync.Mutex
}

type User struct {
	Name       string    `bson:"name"`
	Key        string    `bson:"key"`
	ExpiryDate time.Time `bson:"expiry_date"`
	Solves     int       `bson:"solves"`
}
