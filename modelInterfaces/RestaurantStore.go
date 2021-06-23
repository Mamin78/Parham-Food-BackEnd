package modelInterfaces

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
