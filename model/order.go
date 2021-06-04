package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	UserPhone    string             `json:"user_phone" bson:"user_phone"`

	State       int       `json:"state" bson:"state"`
	Price       float64   `json:"price" bson:"price"`
	TimeStamp   time.Time `json:"timestamp" bson:"timestamp"`
	WaitingTime time.Time `json:"waiting_time" bson:"waiting_time"`

	Foods []FoodOrder `json:"foods" bson:"foods"`
}

type FoodOrder struct {
	FoodID primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`
	Number int                `json:"number" bson:"number"`
}
