package store

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"myapp/model"
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
