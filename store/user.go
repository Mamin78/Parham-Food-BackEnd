package store

import "go.mongodb.org/mongo-driver/mongo"

type UserStore struct {
	db *mongo.Collection
}

func NewUserStore(db *mongo.Collection) *UserStore {
	return &UserStore{
		db: db,
	}
}
