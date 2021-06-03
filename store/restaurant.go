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

func (rs *RestaurantStore) UpdateInformation(managerEmail string, res *model.Restaurant) error {
	newRes := bson.M{"name": res.Name, "area": res.Area, "address": res.Address, "service_area": res.ServiceArea, "start_working_hours": res.StartWorkingHours, "end_working_hours": res.EndWorkingHours, "base_food_price": res.BaseFoodPrice, "base_food_time": res.BaseFoodTime}
	_, err := rs.db.UpdateOne(context.TODO(), bson.M{"email": managerEmail}, bson.M{"$set": newRes})
	return err
}
