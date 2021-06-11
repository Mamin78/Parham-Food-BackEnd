package modelInterfaces

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
	GetAllRestaurants() (*[]model.Restaurant, error)
}

type FoodStore interface {
	CreateFood(food *model.Food) error
	GetAllFoodsOfRestaurant(resName string) (*model.Restaurant, error)
	GetAllFoodsOfRestaurantByID(resID primitive.ObjectID) ([]model.Food, error)
	EnableFoodByID(foodId string) error
	DisableFoodByID(foodId string) error
	GetFoodByID(id string) (*model.Food, error)
	DeleteFoodByID(foodId string) error
	GetAllFoodsByIDs(foods []model.FoodOrder) (*[]model.Food, error)
	AddCommentToFood(commentId primitive.ObjectID, food *model.Food) error
	GetFoodByPrimitiveTypeID(id primitive.ObjectID) (*model.Food, error)
	AddRateToFood(rate model.Rate, food *model.Food) error
	GetFoodsWithSpecificResAndName(resID primitive.ObjectID, foodName string) (*[]model.Food, error)
	GetAllFoods() (*[]model.Food, error)
	GetAllFavoriteFoods(ids []primitive.ObjectID) (*[]model.Food, error)
	GetAllFoodsOfSomeRestaurants(ids []primitive.ObjectID) (*[]model.Food, error)
	GetAllFoodsOfSomeRestaurantsAndSpecificFoodName(ids []primitive.ObjectID, name string) (*[]model.Food, error)
}

type OrderStore interface {
	CreateOrder(order *model.Order) error
	GetAllRestaurantOrdersByIDs(ordersID []primitive.ObjectID) (*[]model.Order, error)
	GetOrderByID(id string) (*model.Order, error)
	ChangeOrderStatus(orderID string, status int) error
	GetAllUserOrders(userID primitive.ObjectID) (*[]model.Order, error)
	ChangeOrderAcceptTime(orderID string, time time.Time) error
}

type CommentStore interface {
	CreateComment(comment *model.Comment) error
	GetAllFoodComments(foodID string) (*[]model.Comment, error)
	GetCommentByID(ID primitive.ObjectID) (*model.Comment, error)
	AddManagerReply(managerReply model.ManagerReply, commentId primitive.ObjectID) error
}

type ManagerCommentStore interface {
}
