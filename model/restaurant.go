package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseRestaurant struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	ManagerEmail string             `json:"manager_email" bson:"manager_email"`

	Name         string `json:"name" bson:"name"`
	WorkingHours string `json:"working_hours" bson:"working_hours"`
	Region       string `json:"region" bson:"region"`
	Address      string `json:"address" bson:"address"`

	Token string `json:"token,omitempty" bson:"-"`

	Orders *[]primitive.ObjectID `json:"orders" bson:"orders"`
}

type Restaurant struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`

	Email    string `json:"email,omitempty" bson:"email"`
	Password string `json:"password" bson:"password"`

	Name              string  `json:"name" bson:"name"`
	Area              int     `json:"area" bson:"area"`
	Address           string  `json:"address" bson:"address"`
	ServiceArea       []int   `json:"service_area" bson:"service_area"`
	StartWorkingHours string  `json:"start_working_hours" bson:"start_working_hours"`
	EndWorkingHours   string  `json:"end_working_hours" bson:"end_working_hours"`
	BaseFoodPrice     float64 `json:"base_food_price" bson:"base_food_price"`
	BaseFoodTime      float64 `json:"base_food_time" bson:"base_food_time"`

	Foods  *[]primitive.ObjectID `json:"foods" bson:"foods"`
	Orders *[]primitive.ObjectID `json:"orders" bson:"orders"`
}

func NewRestaurant(res *Restaurant) *Restaurant {
	r := new(Restaurant)
	r.ID = res.ID
	r.Email = res.Email
	r.Name = res.Name
	r.Area = res.Area
	r.Address = res.Address
	r.ServiceArea = res.ServiceArea
	r.StartWorkingHours = res.StartWorkingHours
	r.EndWorkingHours = res.EndWorkingHours
	r.BaseFoodTime = res.BaseFoodTime
	r.BaseFoodPrice = res.BaseFoodPrice
	return r
}
