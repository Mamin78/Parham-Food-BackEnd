package store

import "go.mongodb.org/mongo-driver/mongo"

type OrderStore struct {
	db *mongo.Collection
}

func NewOrderStore(db *mongo.Collection) *OrderStore {
	return &OrderStore{
		db: db,
	}
}

