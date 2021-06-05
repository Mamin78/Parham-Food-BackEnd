package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	//RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	FoodID primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id"`

	Text string `json:"text" bson:"text"`
	//Rate int    `json:"rate" bson:"rate"`
}

type ManagerComment struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	//ParentID     primitive.ObjectID `json:"parent_id,omitempty" bson:"parent_id"` //must be unique!
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	FoodID       primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`

	Text string `json:"text" bson:"text"`
}

//should i create a struct for Rate ?
