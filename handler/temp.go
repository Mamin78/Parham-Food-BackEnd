package handler

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/model"
	"myapp/utils"
)

type restaurantResponse struct {
	Restaurant struct {
		ID    primitive.ObjectID `json:"id,omitempty" bson:"_id"`
		Email string             `json:"email,omitempty" bson:"email"`
		Token string             `json:"token"`
	} `json:"user"`
}

func newRestaurantResponse(u *model.Restaurant) *restaurantResponse {
	r := new(restaurantResponse)
	r.Restaurant.ID = u.ID
	r.Restaurant.Email = u.Email
	r.Restaurant.Token = utils.GenerateJWT(u.ID.String())
	return r
}

