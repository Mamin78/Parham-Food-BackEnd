package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/Mamin78/Parham-Food-BackEnd/utils"
)

type managerResponse struct {
	Manager struct {
		Email string `json:"email,omitempty" bson:"email"`
		Token string `json:"token"`
	} `json:"user"`
}

func newManagerResponse(u *model.BaseManager) *managerResponse {
	r := new(managerResponse)
	r.Manager.Email = u.Email
	r.Manager.Token = utils.GenerateJWT(u.Email)
	return r
}

type userResponse struct {
	User struct {
		PhoneNumber string `json:"phone_number,omitempty" bson:"phone_number"`
		Token       string `json:"token"`
	} `json:"user"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.PhoneNumber = u.PhoneNumber
	r.User.Token = utils.GenerateUserJWT(u.PhoneNumber)
	return r
}
