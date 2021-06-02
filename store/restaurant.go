package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"myapp/model"
)

type RestaurantStore struct {
	db *mongo.Collection
}

func NewRestaurantStore(db *mongo.Collection) *RestaurantStore {
	return &RestaurantStore{
		db: db,
	}
}

func (rs *RestaurantStore) CreateRestaurant(restaurant *model.Restaurant) error {
	_, err := rs.db.InsertOne(context.TODO(), restaurant)
	return err
}

func (rs *RestaurantStore) GetRestaurantByManagerEmail(email string) (*model.Restaurant, error) {
	var u model.Restaurant
	err := rs.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&u)
	return &u, err
}
