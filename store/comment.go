package store

import (
	"context"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentStore struct {
	db *mongo.Collection
}

func NewCommentStore(db *mongo.Collection) *CommentStore {
	return &CommentStore{
		db: db,
	}
}

func (cs *CommentStore) CreateComment(comment *model.Comment) error {
	_, err := cs.db.InsertOne(context.TODO(), comment)
	return err
}
