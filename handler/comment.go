package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

func (h *Handler) AddUserCommentToFood(c echo.Context) error {
	userPhone := stringFieldFromToken(c, "phone")

	userComment := new(model.Comment)
	userComment.ID = primitive.NewObjectID()

	//here, we have restaurant id and food id !
	//the costumer had to buy this food !
	if err := c.Bind(&userComment); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	user, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	userComment.UserID = user.ID

	orders, err := h.ordersStore.GetAllUserOrders(user.ID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	if !FoodIsInOrders(orders, userComment.FoodID) {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "you have not bought this food", false))
	}

	food, err := h.foodsStore.GetFoodByPrimitiveTypeID(userComment.FoodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "food not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.foodsStore.AddCommentToFood(userComment.ID, food)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "food not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.userStore.AddCommentToUser(userComment.ID, user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.commentsStore.CreateComment(userComment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(userComment, "", true))
}

func FoodIsInOrders(orders *[]model.Order, foodID primitive.ObjectID) bool {
	for _, order := range *orders {
		for _, food := range order.Foods {
			if foodID == food.FoodID {
				return true
			}
		}
	}
	return false
}
