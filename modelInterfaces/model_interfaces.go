package modelInterfaces

import "myapp/model"

type UserStore interface {
	CreateUser(restaurant *model.User) error
	GetUserByPhone(email string) (*model.User, error)
}

type RestaurantStore interface {
	CreateRestaurant(restaurant *model.Restaurant) error
	UpdateInformation(managerEmail string, res *model.Restaurant) error
	GetRestaurantByManagerEmail(email string) (*model.Restaurant, error)
	GetRestaurantById(id string) (*model.Restaurant, error)
}

type FoodStore interface {
}

type OrderStore interface {
}

type CommentStore interface {
}

type ManagerCommentStore interface {
}
