package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
	UserID       primitive.ObjectID `json:"user_id,omitempty" bson:"user_id"`

	State int     `json:"state" bson:"state"`
	Price float64 `json:"price" bson:"price"`

	//TimeStamp   time.Time `json:"timestamp" bson:"timestamp"`

	AcceptTime       time.Time     `json:"accept_time" bson:"accept_time"`
	PreparationTime  time.Duration `json:"preparation_time" bson:"preparation_time"`
	TransmissionTime time.Duration `json:"transmission_time" bson:"transmission_time"`
	//RemainingTime    time.Time     `json:"remaining_time" bson:"remaining_time"`

	Foods []FoodOrder `json:"foods" bson:"foods"`
}

type FoodOrder struct {
	FoodID primitive.ObjectID `json:"food_id,omitempty" bson:"food_id"`
	Number int                `json:"number" bson:"number"`
}

type OrderTimes struct {
	AcceptTime       time.Time `json:"accept_time" bson:"accept_time"`
	PreparationTime  time.Time `json:"reparation_time" bson:"reparation_time"`
	TransmissionTime time.Time `json:"transmission_time" bson:"transmission_time"`
	WaitingTime      time.Time `json:"waiting_time" bson:"waiting_time"`
}
