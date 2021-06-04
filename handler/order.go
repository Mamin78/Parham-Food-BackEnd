package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"myapp/model"
	"net/http"
)

func (h *Handler) CreateOrder(c echo.Context) error {
	userPhone := stringFieldFromToken(c, "phone")

	fmt.Println(userPhone)
	userOrder := new(model.Order)
	userOrder.ID = primitive.NewObjectID()
	//food.RestaurantName = res.Name
	//food.RestaurantID = res.ID
	if err := c.Bind(&userOrder); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	foods, err := h.foodsStore.GetAllFoodsByIDs(userOrder.Foods)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	isOne := IsFromOneRestaurant(foods)
	if !isOne {
		return c.JSON(http.StatusBadRequest, "foods are not from one restaurant")
	}
	//res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	//if err != nil {
	//	if err == mgo.ErrNotFound {
	//		return c.JSON(http.StatusUnauthorized, "invalid Restaurant!")
	//	}
	//	return c.JSON(http.StatusBadRequest, "Bad Request")
	//}
	//res.Password = ""
	//
	//food := new(model.Food)
	//food.ID = primitive.NewObjectID()
	////food.RestaurantName = res.Name
	//food.RestaurantID = res.ID
	//if err := c.Bind(&food); err != nil {
	//	return c.JSON(http.StatusBadRequest, "Bad Request")
	//}
	//err = h.foodsStore.CreateRestaurant(food)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, "Bad Request")
	//}
	//
	//err = h.restaurantStore.AddFoodToRestaurant(res.Name, food, res)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, "Bad Request")
	//}
	return c.JSON(http.StatusCreated, foods)
}

func IsFromOneRestaurant(foods *[]model.Food) bool {
	set := make(map[primitive.ObjectID]bool)
	for _, food := range *foods {
		set[food.RestaurantID] = true
	}
	//for i := 0; i < len(*foods); i++ {
	//	set[foods[i].RestaurantID] = true
	//}
	if len(set) > 1 {
		return false
	}
	return true
}
