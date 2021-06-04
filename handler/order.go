package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

const (
	managerConfirmationStatus = 1
)

func (h *Handler) CreateOrder(c echo.Context) error {
	userPhone := stringFieldFromToken(c, "phone")

	userOrder := new(model.Order)
	userOrder.ID = primitive.NewObjectID()

	if err := c.Bind(&userOrder); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	foods, err := h.foodsStore.GetAllFoodsByIDs(userOrder.Foods)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid foods", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	if !IsAllFoodSEnable(foods) {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "some foods are not enable!", false))
	}

	resID, isOne := IsFromOneRestaurant(foods)
	if !isOne {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "foods are not from one restaurant", false))
	}

	user, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid foods", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	cost := calculateOrderPrice(foods, userOrder.Foods)
	if cost > user.Credit {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "sorry, your credit is not sufficient", false))
	}

	user.Credit -= cost

	err = h.userStore.AddOrderToUserByID(userOrder, user)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "user not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.userStore.UpdateUserCredit(user.ID, user.Credit)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user!", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	res, err := h.restaurantStore.GetRestaurantByPrimitiveTypeId(resID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "restaurant not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.restaurantStore.AddOrderToRestaurantByID(userOrder, res)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "restaurant not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	userOrder.RestaurantID = res.ID
	userOrder.UserID = user.ID
	err = h.ordersStore.CreateOrder(userOrder)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(userOrder, "", true))
}

func (h *Handler) GetRestaurantOrders(c echo.Context) (err error) {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Manager", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	res.Password = ""

	orders, err := h.ordersStore.GetAllRestaurantOrdersByIDs(res.Orders)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "restaurant not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(orders, "", true))
}

func (h *Handler) ConfirmOrderByRestaurantManager(c echo.Context) (err error) {
	managerEmail := stringFieldFromToken(c, "email")
	orderID := c.Param("order_id")

	order, err := h.ordersStore.GetOrderByID(orderID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "order not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "restaurant not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	res.Password = ""
	if res.ID != order.RestaurantID {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "access denied", false))

	}

	err = h.ordersStore.ChangeOrderStatus(orderID, managerConfirmationStatus)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "order not found", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse("This order confirmed by you.", "", true))
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

func IsAllFoodSEnable(foods *[]model.Food) bool {
	for _, food := range *foods {
		if !food.CanBeOrdered {
			return false
		}
	}
	return true
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
