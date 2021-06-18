package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ManagerReply struct {
	Text string `json:"text" bson:"text"`
}

type Comment struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	//RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	FoodID primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id"`

	Text  string       `json:"text" bson:"text"`
	Reply ManagerReply `json:"reply" bson:"reply"`
}

type ManagerComment struct {
	ParentID     primitive.ObjectID `json:"parent_id,omitempty" bson:"parent_id"` //must be unique!
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	//FoodID       primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`

	Text string `json:"text" bson:"text"`
}

type CommentAsResponse struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	FoodID primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`

	UserID   primitive.ObjectID `json:"user_id,omitempty" bson:"user_id"`
	UserName string             `json:"user_name,omitempty" bson:"user_name"`

	Text  string       `json:"text" bson:"text"`
	Reply ManagerReply `json:"reply" bson:"reply"`
}

func CreateRespCommentFromComment(comment Comment) *CommentAsResponse {
	r := new(CommentAsResponse)
	r.ID = comment.ID
	r.FoodID = comment.FoodID
	r.UserID = comment.UserID
	r.Text = comment.Text
	r.Reply = comment.Reply
	return r
}
