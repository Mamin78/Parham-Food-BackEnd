package store

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"myapp/model"
)

type FoodStore struct {
	db *mongo.Collection
}

func NewFoodStore(db *mongo.Collection) *FoodStore {
	return &FoodStore{
		db: db,
	}
}

func (fs *FoodStore) CreateRestaurant(food *model.Food) error {
	_, err := fs.db.InsertOne(context.TODO(), food)
	return err
}

func (fs *FoodStore) GetAllFoodsOfRestaurant(resName string) (*model.Restaurant, error) {
	var u model.Restaurant
	var uu []model.Food
	//query := bson.M{"_id": bson.M{"$in": usernames}}
	query := bson.M{"restaurant_name": resName}
	foods, err := fs.db.Find(context.TODO(), query)

	if err != nil {
		return nil, err
	}
	if err = foods.All(context.TODO(), &uu); err != nil {
		return nil, err
	}
	fmt.Println("foods")
	fmt.Println(uu)
	fmt.Println("foods")
	return &u, err
}

func (fs *FoodStore) GetAllFoodsOfRestaurantByID(resID primitive.ObjectID) ([]model.Food, error) {
	var uu []model.Food
	//query := bson.M{"_id": bson.M{"$in": usernames}}
	query := bson.M{"restaurant_id": resID}
	foods, err := fs.db.Find(context.TODO(), query)

	if err != nil {
		return nil, err
	}
	if err = foods.All(context.TODO(), &uu); err != nil {
		return nil, err
	}
	fmt.Println("foods")
	fmt.Println(uu)
	fmt.Println("foods")
	return uu, err
}
