package store

import (
	"fmt"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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

func (fs *FoodStore) DisableFoodByID(foodId string) error {
	oid, err := primitive.ObjectIDFromHex(foodId)
	if err != nil {
		return nil
	}
	newRes := bson.M{"can_be_ordered": false}
	_, err = fs.db.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": newRes})
	return err
}

func (fs *FoodStore) EnableFoodByID(foodId string) error {
	oid, err := primitive.ObjectIDFromHex(foodId)
	if err != nil {
		return nil
	}
	newRes := bson.M{"can_be_ordered": true}
	_, err = fs.db.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": newRes})
	return err
}

func (fs *FoodStore) DeleteFoodByID(foodId string) error {
	oid, err := primitive.ObjectIDFromHex(foodId)
	if err != nil {
		return nil
	}
	_, err = fs.db.DeleteOne(context.TODO(), bson.M{"_id": oid})
	return err
}

func (fs *FoodStore) GetFoodByID(id string) (*model.Food, error) {
	var u model.Food
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &u, nil
	}
	err = fs.db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return &u, err
}

func (fs *FoodStore) GetAllFoodsByIDs(foods []model.FoodOrder) (*[]model.Food, error) {
	var ids []primitive.ObjectID
	for _, food := range foods {
		ids = append(ids, food.FoodID)
	}
	var result []model.Food
	query := bson.M{"_id": bson.M{"$in": ids}}
	res, err := fs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return &result, err
}
