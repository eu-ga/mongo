package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var db *mongo.Client

//CreateClient creates client for mongodb server on uri+port and returns a  pointer and an error
func CreateMongoClient(uri, port string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+uri+port))
	if err != nil {
		return nil, errors.New("cannot connect to the mongo: " + err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.New("cannot ping mongo: " + err.Error())
	}
	log.Println("Connected to MongoDB!")
	db = client
	return client, nil
}

//CloseClient checks if mongodb client if pingable and disconnects from database
func CloseClient(ctx context.Context) error {
	err := db.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.New("cannot ping mongo: " + err.Error())
	}
	return db.Disconnect(ctx)
}

//GetCollection gets collection with a collectionName from database
func GetCollection(database, collectionName string) *mongo.Collection {
	return db.Database(database).Collection(collectionName)
}
