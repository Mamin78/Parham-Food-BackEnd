package db

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DataBaseName = "parham-food"
)

func SetupUsersDb(mongoClient *mongo.Client) *mongo.Collection {
	usersDb := mongoClient.Database(DataBaseName).Collection("users")
	createUniqueIndices(usersDb, "phone_number")
	return usersDb
}

func SetupRestaurantsDb(mongoClient *mongo.Client) *mongo.Collection {
	restaurantsDb := mongoClient.Database(DataBaseName).Collection("restaurants")
	createUniqueIndices(restaurantsDb, "email")
	return restaurantsDb
}

func SetupFoodsDb(mongoClient *mongo.Client) *mongo.Collection {
	foodsDb := mongoClient.Database(DataBaseName).Collection("foods")
	return foodsDb
}

func SetupOrdersDb(mongoClient *mongo.Client) *mongo.Collection {
	ordersDb := mongoClient.Database(DataBaseName).Collection("orders")
	return ordersDb
}

func SetupCommentsDb(mongoClient *mongo.Client) *mongo.Collection {
	commentsDb := mongoClient.Database(DataBaseName).Collection("comments")
	return commentsDb
}

func SetupManagerCommentsDb(mongoClient *mongo.Client) *mongo.Collection {
	managerCommentsDb := mongoClient.Database(DataBaseName).Collection("manager-comments")
	return managerCommentsDb
}

func createUniqueIndices(db *mongo.Collection, field string) {
	_, err := db.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: field, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

//func SetupTweetsDb(mongoClient *mongo.Client) *mongo.Collection {
//	tweetsDb := mongoClient.Database("parham-food").Collection("tweets")
//	_, err := tweetsDb.Indexes().CreateOne(
//		context.Background(),
//		mongo.IndexModel{
//			Keys:    bson.D{{Key: "text", Value: "text"}},
//			Options: nil,
//		})
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	return tweetsDb
//}
