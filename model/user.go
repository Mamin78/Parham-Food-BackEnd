package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`

	Name        string  `json:"name" bson:"name"`
	PhoneNumber string  `json:"phone_number,omitempty" bson:"phone_number"`
	Password    string  `json:"password" bson:"password"`
	Credit      float64 `json:"credit" bson:"credit"`
	Area        int     `json:"area" bson:"area"`
	Address     string  `json:"address" bson:"address"`

	Orders *[]primitive.ObjectID `json:"orders" bson:"orders"`
}
