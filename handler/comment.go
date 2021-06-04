package handler

import (
	"fmt"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

func (h *Handler) AddUserCommentToFood(c echo.Context) error {
	userPhone := stringFieldFromToken(c, "phone")
	foodID := c.Param("food_id")

	fmt.Println(foodID)
	userComment := new(model.Comment)
	userComment.ID = primitive.NewObjectID()

	if err := c.Bind(&userComment); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request1", false))
	}

	user, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request2", false))
	}

	orders, err := h.ordersStore.GetAllUserOrders(user.ID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request3", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(orders, "", true))
}
