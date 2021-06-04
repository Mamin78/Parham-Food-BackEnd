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

	if err := c.Bind(&userOrder); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	foods, err := h.foodsStore.GetAllFoodsByIDs(userOrder.Foods)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid foods")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	resID, isOne := IsFromOneRestaurant(foods)
	fmt.Println("resID")
	fmt.Println(resID)
	fmt.Println("resID")
	if !isOne {
		return c.JSON(http.StatusBadRequest, "foods are not from one restaurant")
	}

	user, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid foods")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	cost := calculateOrderPrice(foods, userOrder.Foods)
	if cost > user.Credit {
		return c.JSON(http.StatusBadRequest, "sorry, your credit is not sufficient")
	}

	user.Credit -= cost

	err = h.userStore.AddOrderToUserByID(userOrder, user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid foods")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	err = h.userStore.UpdateUserCredit(user.ID, user.Credit)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid user!")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	fmt.Println("resID")
	fmt.Println(resID)
	fmt.Println("resID")
	res, err := h.restaurantStore.GetRestaurantById(resID.String())
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid foods")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	fmt.Println("res")
	fmt.Println(res)
	fmt.Println("res")

	err = h.restaurantStore.AddOrderToRestaurantByID(userOrder, res)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid foods")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	userOrder.RestaurantID = res.ID
	fmt.Println("userOrder.RestaurantID")
	fmt.Println(userOrder.RestaurantID)
	fmt.Println("userOrder.RestaurantID")
	userOrder.UserID = user.ID
	err = h.ordersStore.CreateOrder(userOrder)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid foods")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, userOrder)
}

func IsFromOneRestaurant(foods *[]model.Food) (primitive.ObjectID, bool) {
	set := make(map[primitive.ObjectID]bool)
	var no primitive.ObjectID
	for _, food := range *foods {
		set[food.RestaurantID] = true
	}
	if len(set) > 1 {
		return no, false
	}

	for k, _ := range set {
		return k, true
	}
	return no, false
}

func calculateOrderPrice(foods *[]model.Food, orderFoods []model.FoodOrder) float64 {
	price := make(map[primitive.ObjectID]float64)
	for _, food := range *foods {
		price[food.ID] = food.Price
	}

	cost := 0.0
	for _, food := range orderFoods {
		cost += float64(food.Number) * price[food.FoodID]
	}
	return cost
}
