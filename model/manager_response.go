package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ManagerResponse struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	CommentID    primitive.ObjectID `json:"comment_id,omitempty" bson:"comment_id"`
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	FoodID       primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`
	ManagerEmail string             `json:"manager_email" bson:"manager_email"`

	Text string `json:"text" bson:"text"`
}
