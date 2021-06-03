package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (rs *RestaurantStore) GetRestaurantById(id string) (*model.Restaurant, error) {
	var u model.Restaurant
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &u, nil
	}
	err = rs.db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return &u, err
}

func (rs *RestaurantStore) GetRestaurantByName(name string) (*model.Restaurant, error) {
	var u model.Restaurant
	err := rs.db.FindOne(context.TODO(), bson.M{"name": name}).Decode(&u)
	return &u, err
}

func (rs *RestaurantStore) UpdateInformation(managerEmail string, res *model.Restaurant) error {
	newRes := bson.M{"name": res.Name, "area": res.Area, "address": res.Address, "service_area": res.ServiceArea, "start_working_hours": res.StartWorkingHours, "end_working_hours": res.EndWorkingHours, "base_food_price": res.BaseFoodPrice, "base_food_time": res.BaseFoodTime}
	_, err := rs.db.UpdateOne(context.TODO(), bson.M{"email": managerEmail}, bson.M{"$set": newRes})
	return err
}

func (rs *RestaurantStore) AddFoodToRestaurant(resName string, food *model.Food, res *model.Restaurant) error {
	res.Foods = append(res.Foods, food.ID)
	newRes := bson.M{"foods": res.Foods}
	_, err := rs.db.UpdateOne(context.TODO(), bson.M{"name": resName}, bson.M{"$set": newRes})
	return err
}

