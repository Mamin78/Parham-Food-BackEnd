package modelInterfaces

import "myapp/model"

type UserStore interface {
}

type RestaurantStore interface {
	CreateRestaurant(restaurant *model.Restaurant) error
	GetRestaurantByManagerEmail(email string) (*model.Restaurant, error)
}

type FoodStore interface {
}

type OrderStore interface {
}

type CommentStore interface {
}

type ManagerCommentStore interface {
}
