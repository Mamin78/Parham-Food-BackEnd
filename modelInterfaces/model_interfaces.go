package modelInterfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/model"
)

type UserStore interface {
	CreateUser(restaurant *model.User) error
	GetUserByPhone(email string) (*model.User, error)
	UpdateUserInfoByPhone(phoneNumber string, user *model.User) error
}

type RestaurantStore interface {
	CreateRestaurant(restaurant *model.Restaurant) error
	UpdateInformation(managerEmail string, res *model.Restaurant) error
	GetRestaurantByManagerEmail(email string) (*model.Restaurant, error)
	GetRestaurantById(id string) (*model.Restaurant, error)
	GetRestaurantByName(name string) (*model.Restaurant, error)
	AddFoodToRestaurant(resName string, food *model.Food, res *model.Restaurant) error
	DeleteFoodFromRestaurant(foodID string, res *model.Restaurant) error
}

type FoodStore interface {
	CreateRestaurant(food *model.Food) error
	GetAllFoodsOfRestaurant(resName string) (*model.Restaurant, error)
	GetAllFoodsOfRestaurantByID(resID primitive.ObjectID) ([]model.Food, error)
	EnableFoodByID(foodId string) error
	DisableFoodByID(foodId string) error
	GetFoodByID(id string) (*model.Food, error)
	DeleteFoodByID(foodId string) error
}

type OrderStore interface {
}

type CommentStore interface {
}

type ManagerCommentStore interface {
}
