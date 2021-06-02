package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"myapp/modelInterfaces"
	"net/http"
)

type (
	Handler struct {
		userStore            modelInterfaces.UserStore
		restaurantStore      modelInterfaces.RestaurantStore
		foodsStore           modelInterfaces.FoodStore
		ordersStore          modelInterfaces.OrderStore
		commentsStore        modelInterfaces.CommentStore
		managerCommentsStore modelInterfaces.ManagerCommentStore
	}
)

func NewHandler(userStore modelInterfaces.UserStore, restaurantStore modelInterfaces.RestaurantStore, foodsStore modelInterfaces.FoodStore, ordersStore modelInterfaces.OrderStore, commentsStore modelInterfaces.CommentStore, managerCommentsStore modelInterfaces.ManagerCommentStore) (handler *Handler) {
	return &Handler{
		userStore:            userStore,
		restaurantStore:      restaurantStore,
		foodsStore:           foodsStore,
		ordersStore:          ordersStore,
		commentsStore:        commentsStore,
		managerCommentsStore: managerCommentsStore,
	}
}

func NewHandlerNotPointer(userStore modelInterfaces.UserStore, restaurantStore modelInterfaces.RestaurantStore, foodsStore modelInterfaces.FoodStore, ordersStore modelInterfaces.OrderStore, commentsStore modelInterfaces.CommentStore, managerCommentsStore modelInterfaces.ManagerCommentStore) (handler *Handler) {
	var h *Handler
	h.userStore = userStore
	h.restaurantStore = restaurantStore
	h.foodsStore = foodsStore
	h.ordersStore = ordersStore
	h.commentsStore = commentsStore
	h.managerCommentsStore = managerCommentsStore
	return h
}

func (h *Handler) BaseRouter(c echo.Context) error {
	return c.JSON(http.StatusCreated, errors.New("welcome to parhamfood"))
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
