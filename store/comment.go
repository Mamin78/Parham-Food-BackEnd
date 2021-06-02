package store

import "go.mongodb.org/mongo-driver/mongo"

type CommentStore struct {
	db *mongo.Collection
}

func NewCommentStore(db *mongo.Collection) *CommentStore {
	return &CommentStore{
		db: db,
	}
}
