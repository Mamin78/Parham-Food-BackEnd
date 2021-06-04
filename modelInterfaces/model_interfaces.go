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

	//AddOrderToUserByPhone(phone string, newOrder *model.Order, user *model.User) error
	AddOrderToUserByID(newOrder *model.Order, user *model.User) error
}

type RestaurantStore interface {
	CreateRestaurant(restaurant *model.Restaurant) error
	UpdateInformation(managerEmail string, res *model.Restaurant) error
	GetRestaurantByManagerEmail(email string) (*model.Restaurant, error)
	GetRestaurantById(id string) (*model.Restaurant, error)
	GetRestaurantByName(name string) (*model.Restaurant, error)
	AddFoodToRestaurant(resName string, food *model.Food, res *model.Restaurant) error
	DeleteFoodFromRestaurant(foodID string, res *model.Restaurant) error
	AddOrderToRestaurantByID(newOrder *model.Order, res *model.Restaurant) error
	GetRestaurantByPrimitiveTypeId(id primitive.ObjectID) (*model.Restaurant, error)
}

type FoodStore interface {
	CreateRestaurant(food *model.Food) error
	GetAllFoodsOfRestaurant(resName string) (*model.Restaurant, error)
	GetAllFoodsOfRestaurantByID(resID primitive.ObjectID) ([]model.Food, error)
	EnableFoodByID(foodId string) error
	DisableFoodByID(foodId string) error
	GetFoodByID(id string) (*model.Food, error)
	DeleteFoodByID(foodId string) error
	GetAllFoodsByIDs(foods []model.FoodOrder) (*[]model.Food, error)
}

type OrderStore interface {
	CreateOrder(order *model.Order) error
	GetAllRestaurantOrdersByIDs(ordersID []primitive.ObjectID) (*[]model.Order, error)
	GetOrderByID(id string) (*model.Order, error)
	ChangeOrderStatus(orderID string, status int) error
}

type CommentStore interface {
}

type ManagerCommentStore interface {
}
