package store

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type OrderStore struct {
	db *mongo.Collection
}

func NewOrderStore(db *mongo.Collection) *OrderStore {
	return &OrderStore{
		db: db,
	}
}

func (orderStore *OrderStore) CreateOrder(order *model.Order) error {
	_, err := orderStore.db.InsertOne(context.TODO(), order)
	return err
}

func (orderStore *OrderStore) GetAllRestaurantOrdersByIDs(ordersID []primitive.ObjectID) (*[]model.Order, error) {
	var result []model.Order
	query := bson.M{"_id": bson.M{"$in": ordersID}}
	res, err := orderStore.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return &result, err
}

func (orderStore *OrderStore) GetOrderByID(id string) (*model.Order, error) {
	var u model.Order
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &u, nil
	}
	err = orderStore.db.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	return &u, err
}

func (orderStore *OrderStore) ChangeOrderStatus(orderID string, status int) error {
	oid, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil
	}
	newRes := bson.M{"state": status}
	_, err = orderStore.db.UpdateOne(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": newRes})
	return err
}

func (orderStore *OrderStore) GetAllUserOrders(userID primitive.ObjectID) (*[]model.Order, error) {
	var result []model.Order
	query := bson.M{"user_id": userID}
	res, err := orderStore.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return &result, err
}
