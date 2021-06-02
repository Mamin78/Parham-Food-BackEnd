package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`

	Name         string  `json:"name" bson:"name"`
	CanBeOrdered bool    `json:"can _be_ordered" bson:"can _be_ordered"`
	Rate         float64 `json:"rate" bson:"rate"`
	Price        float64 `json:"price" bson:"price"`
}
