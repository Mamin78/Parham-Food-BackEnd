package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseRestaurant struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	ManagerEmail string             `json:"manager_email" bson:"manager_email"`

	Name         string  `json:"name" bson:"name"`
	WorkingHours string  `json:"working-hours" bson:"working_hours"`
	Region       string  `json:"region" bson:"region"`
	Address      string  `json:"address" bson:"address"`

	Token string `json:"token,omitempty" bson:"-"`

	Orders *[]primitive.ObjectID `json:"orders" bson:"orders"`
}

type Restaurant struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`

	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password" bson:"password"`

	Name         string  `json:"name" bson:"name"`
	WorkingHours string  `json:"working-hours" bson:"working_hours"`
	Region       string  `json:"region" bson:"region"`
	Address      string  `json:"address" bson:"address"`

	Token string `json:"token" bson:"-"`

	Foods  *[]primitive.ObjectID `json:"foods" bson:"foods"`
	Orders *[]primitive.ObjectID `json:"orders" bson:"orders"`
}
