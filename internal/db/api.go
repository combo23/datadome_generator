package db

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDBInstance(uri string) *Database {
	return &Database{MongoURI: uri}
}

func (t *Database) Connect() {
	t.mu.Lock()
	defer t.mu.Unlock()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	t.client = client
}

func (t *Database) UpdateUsage(apikey string) error {
	filter := bson.M{"key": apikey}

	// Define an update to decrement the "solves" field by 1
	update := bson.M{
		"$inc": bson.M{"solves": -1},
	}

	// Perform the update
	_, err := t.client.Database("users").Collection("customers").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (t *Database) BanUser(apikey string) error {
	filter := bson.M{"key": apikey}

	// Define an update to set the "solves" field to -1
	update := bson.M{
		"$set": bson.M{"solves": -1},
	}

	// Perform the update
	_, err := t.client.Database("users").Collection("customers").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (t *Database) GetSolves(apikey string) (int, error) {
	filter := bson.M{"key": apikey}

	result := t.client.Database("users").Collection("customers").FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return 0, result.Err()
	}

	var user User
	if err := result.Decode(&user); err != nil {
		return 0, err
	}
	return user.Solves, nil
}

func (t *Database) ValidateAPIKey(apikey string) (bool, error) {

	if apikey == "" {
		return false, ErrNoAPIKey
	}

	if !isUUID(apikey) {
		return false, ErrInvalidAPIKey
	}

	filter := bson.M{"key": apikey}

	// Create a context with a timeout

	// Find the document in the "users" collection based on the filter
	result := t.client.Database("users").Collection("customers").FindOne(context.TODO(), filter)
	if result.Err() != nil {
		return false, ErrInvalidAPIKey
	}

	// Decode the document into a User struct
	var user User
	if err := result.Decode(&user); err != nil {
		return false, ErrInternalError
	}
	if user.Name == "combo" {
		return true, nil //exception for me lol
	}
	if user.Solves < 1 || user.ExpiryDate.Before(time.Now()) {
		return false, ErrNoSolves //invalid
	}
	return true, nil //valid
}

func (t *Database) AddUser(name string, solves int) (string, error) {
	apikey := uuid.New().String()
	_, err := t.client.Database("users").Collection("customers").InsertOne(context.TODO(), bson.D{
		{Key: "key", Value: apikey},
		{Key: "solves", Value: solves},
		{Key: "expiry_date", Value: time.Now().AddDate(0, 0, 30)},
		{Key: "name", Value: name},
	})
	if err != nil {
		return "", err
	}
	return apikey, nil
}
