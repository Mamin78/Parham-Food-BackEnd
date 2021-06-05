package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

func (h *Handler) AddUserRateToFood(c echo.Context) error {
	userPhone := stringFieldFromToken(c, "phone")
	foodID := c.Param("food_id")

	userRate := new(model.Rate)

	//here, we have restaurant id and food id !
	//the costumer had to buy this food !
	if err := c.Bind(&userRate); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	user, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	userRate.UserID = user.ID

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "food not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	if hasRated(food.Rates,user.ID) {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "you have rated on this food", false))
	}

	err = h.foodsStore.AddRateToFood(*userRate, food)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "food not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(userRate, "", true))
}

func hasRated(foodRates []model.Rate, userId primitive.ObjectID) bool {
	for _, rate := range foodRates {
		if rate.UserID == userId {
			return true
		}
	}
	return false
}
