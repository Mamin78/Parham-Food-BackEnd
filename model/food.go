package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`

	Name         string  `json:"name" bson:"name"`
	CanBeOrdered bool    `json:"can_be_ordered" bson:"can_be_ordered"`
	//Rate         float64 `json:"rate" bson:"rate"`
	Price        float64 `json:"price" bson:"price"`

	Comments []primitive.ObjectID `json:"comments" bson:"comments"`
	Rates    []Rate               `json:"rates" bson:"rates"`
}

type Rate struct {
	UserID primitive.ObjectID `json:"user_id,omitempty" bson:"user_id"`
	Rate   int                `json:"rate" bson:"rate"`
}
