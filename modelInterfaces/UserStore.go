package modelInterfaces

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	CreateUser(user *model.User) error
	GetUserByPhone(phoneNumber string) (*model.User, error)
	UpdateUserInfoByPhone(phoneNumber string, user *model.User) error

	GetUserByID(userID string) (*model.User, error)
	UpdateUserInfoByID(userID string, user *model.User) error
	UpdateUserCredit(userID primitive.ObjectID, newCredit float64) error

	AddOrderToUserByID(newOrder *model.Order, user *model.User) error

	AddCommentToUser(commentId primitive.ObjectID, user *model.User) error

	GetUserByPrimitiveID(userID primitive.ObjectID) (*model.User, error)
}
