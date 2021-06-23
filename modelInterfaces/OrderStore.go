package modelInterfaces

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderStore interface {
	CreateOrder(order *model.Order) error
	GetAllRestaurantOrdersByIDs(ordersID []primitive.ObjectID) (*[]model.Order, error)
	GetOrderByID(id string) (*model.Order, error)
	ChangeOrderStatus(orderID string, status int) error
	GetAllUserOrders(userID primitive.ObjectID) (*[]model.Order, error)
	ChangeOrderAcceptTime(orderID string, time time.Time) error
}
