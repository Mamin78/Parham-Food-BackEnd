package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
<<<<<<< HEAD
	PATH = "mongodb://localhost:27017"
=======
	//PATH = "mongodb://localhost:27017"
	PATH = "mongodb+srv://parham_food:ahaphf0098@cluster0.3ir6g.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
>>>>>>> 7126080 (mongo atlas)
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(PATH)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		} else {
			err = client.Ping(context.TODO(), nil)
			if err != nil {
				clientInstanceError = err
			}
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
