package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	FoodID       primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`
	UserEmail    string             `json:"user_email" bson:"user_email"`
	ManagerEmail string             `json:"manager_email" bson:"manager_email"`

	Text string `json:"text" bson:"text"`
	Rate int    `json:"rate" bson:"rate"`
}
