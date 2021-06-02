package store

import "go.mongodb.org/mongo-driver/mongo"

type FoodStore struct {
	db *mongo.Collection
}

func NewFoodStore(db *mongo.Collection) *FoodStore {
	return &FoodStore{
		db: db,
	}
}
