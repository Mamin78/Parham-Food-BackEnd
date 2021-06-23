package modelInterfaces

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FoodStore interface {
	CreateFood(food *model.Food) error
	GetAllFoodsOfRestaurant(resName string) (*[]model.Food, error)
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
	GetAllFoodsWithSpecificFoodName(name string) (*[]model.Food, error)
	GetAllFoodsOfRestaurantWithSpecificFoodName(name string, resID primitive.ObjectID) (*[]model.Food, error)
	UpdateFoodRate(foodId string, rate float64) error
}
