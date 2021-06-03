package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name        string  `json:"name" bson:"name"`
	PhoneNumber string  `json:"phone_number,omitempty" bson:"phone_number"`
	Password    string  `json:"password" bson:"password"`
	Credit      float64 `json:"credit" bson:"credit"`

	Token string `json:"token,omitempty" bson:"-"`

	Orders *[]primitive.ObjectID `json:"orders" bson:"orders"`
}


type userResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
		Token    string  `json:"token"`
	} `json:"user"`
}

func newUserResponse(u User) *userResponse {
	r := new(userResponse)
	//r.User.Username = u.Username
	//r.User.Email = u.Email
	//r.User.Token = utils.GenerateJWT(u.ID)
	return r
}


