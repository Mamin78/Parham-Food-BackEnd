package store

import "go.mongodb.org/mongo-driver/mongo"

type ManagerCommentStore struct {
	db *mongo.Collection
}

func NewManagerCommentStore(db *mongo.Collection) *ManagerCommentStore {
	return &ManagerCommentStore{
		db: db,
	}
}