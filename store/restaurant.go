package store

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
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
	err = rs.db.FindOne(context.TODO(), bson.M{"id": oid}).Decode(&u)
	return &u, err
}

func (rs *RestaurantStore) GetRestaurantByPrimitiveTypeId(id primitive.ObjectID) (*model.Restaurant, error) {
	var u model.Restaurant
	err := rs.db.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&u)
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

func (rs *RestaurantStore) DeleteFoodFromRestaurant(foodID string, res *model.Restaurant) error {
	oid, err := primitive.ObjectIDFromHex(foodID)
	if err != nil {
		return err
	}

	newFoods := &[]primitive.ObjectID{}
	for _, fid := range res.Foods {
		if fid != oid {
			*newFoods = append(*newFoods, fid)
		}
	}

	_, err = rs.db.UpdateOne(context.TODO(), bson.M{"_id": res.ID}, bson.M{"$set": bson.M{"foods": newFoods}})
	return err
}

func (rs *RestaurantStore) AddOrderToRestaurantByID(newOrder *model.Order, res *model.Restaurant) error {
	res.Orders = append(res.Orders, newOrder.ID)
	newRes := bson.M{"orders": res.Orders}
	_, err := rs.db.UpdateOne(context.TODO(), bson.M{"_id": res.ID}, bson.M{"$set": newRes})
	return err
}

func (rs *RestaurantStore) GetAllRestaurants() (*[]model.Restaurant, error) {
	var result []model.Restaurant
	query := bson.M{}
	res, err := rs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return &result, err
}
