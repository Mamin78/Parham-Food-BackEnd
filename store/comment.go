package store

import (
	"context"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (cs *CommentStore) GetAllFoodComments(foodID string) (*[]model.Comment, error) {
	var result []model.Comment
	oid, err := primitive.ObjectIDFromHex(foodID)
	if err != nil {
		return &result, nil
	}
	query := bson.M{"food_id": oid}
	res, err := cs.db.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	if err = res.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	return &result, err
}

