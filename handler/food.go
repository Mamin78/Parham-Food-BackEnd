package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"myapp/model"
	"net/http"
)

func (h *Handler) CreateFood(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	res.Password = ""

	food := new(model.Food)
	food.ID = primitive.NewObjectID()
	//food.RestaurantName = res.Name
	food.RestaurantID = res.ID
	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	err = h.foodsStore.CreateRestaurant(food)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	err = h.restaurantStore.AddFoodToRestaurant(res.Name, food, res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusCreated, food)
}

func (h *Handler) DisableFood(c echo.Context) error {
	foodID := c.Param("food_id")
	fmt.Println(foodID)

	err := h.foodsStore.DisableFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusCreated, food)
}

func (h *Handler) EnableFood(c echo.Context) error {
	foodID := c.Param("food_id")
	fmt.Println(foodID)

	err := h.foodsStore.EnableFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusCreated, food)
}

func (h *Handler) DeleteFood(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	res.Password = ""

	foodID := c.Param("food_id")
	fmt.Println(foodID)

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	err = h.foodsStore.DeleteFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	err = h.restaurantStore.DeleteFoodFromRestaurant(foodID,res)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, food)
}
