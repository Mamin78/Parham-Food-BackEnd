package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseManager struct {
	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password" bson:"password"`
}


type Manager struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`

	RestaurantID primitive.ObjectID `json:"restaurant_id,omitempty" bson:"restaurant_id"`
}
