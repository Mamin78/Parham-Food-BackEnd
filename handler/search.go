package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

func (h *Handler) Search(c echo.Context) (err error) {
	//userPhone := stringFieldFromToken(c, "phone")
	query := new(model.Search)
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return doQuery(h, c, query)

	//return c.JSON(http.StatusCreated, model.NewResponse(nil, "", true))
}

func doQuery(h *Handler, c echo.Context, query *model.Search) (err error) {
	if query.Area > 0 && query.Area < 23 {
		if query.RestaurantName != "" {
			res, err := h.restaurantStore.GetRestaurantByName(query.RestaurantName)
			if err != nil {
				if err == mgo.ErrNotFound {
					return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
				}
				return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
			}
			if isInServiceAreas(res, query.Area) {
				if query.FoodName != "" {
					foods, err := h.foodsStore.GetFoodsWithSpecificResAndName(res.ID, query.FoodName)
					if err != nil {
						if err == mgo.ErrNotFound {
							return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
						}
						return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
					}
					return c.JSON(http.StatusOK, model.NewResponse(foods, "", true))
				} else {
					foods, err := h.foodsStore.GetAllFoodsOfRestaurant(query.RestaurantName)
					if err != nil {
						if err == mgo.ErrNotFound {
							return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
						}
						return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
					}
					return c.JSON(http.StatusOK, model.NewResponse(foods, "", true))
				}
			} else {
				return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "Sorry, this restaurant cant service to you at this area", false))
			}
		} else {
			restaurants, err := h.restaurantStore.GetAllRestaurants()
			if err != nil {
				if err == mgo.ErrNotFound {
					return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
				}
				return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
			}

			IDs := allRestaurantsWithSpecificServiceArea(restaurants, query.Area)

			if query.FoodName == "" {
				foods, err := h.foodsStore.GetAllFoodsOfSomeRestaurants(IDs)
				if err != nil {
					if err == mgo.ErrNotFound {
						return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
					}
					return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
				}
				return c.JSON(http.StatusOK, model.NewResponse(foods, "", true))
			} else {
				foods, err := h.foodsStore.GetAllFoodsOfSomeRestaurantsAndSpecificFoodName(IDs, query.FoodName)
				if err != nil {
					if err == mgo.ErrNotFound {
						return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
					}
					return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
				}
				return c.JSON(http.StatusOK, model.NewResponse(foods, "", true))
			}
		}
	} else {
		if query.RestaurantName != "" {

		} else {

		}
	}
	return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
}

func (h *Handler) temp(c echo.Context) (err error) {
	restaurants, err := h.restaurantStore.GetAllRestaurants()
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	IDs := allRestaurantsWithSpecificServiceArea(restaurants, 5)

	foods, err := h.foodsStore.GetAllFoodsOfSomeRestaurants(IDs)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusOK, model.NewResponse(foods, "", true))
}

func isInServiceAreas(res *model.Restaurant, area int) bool {
	for _, serviceArea := range res.ServiceArea {
		if serviceArea == area {
			return true
		}
	}
	return false
}

func allRestaurantsWithSpecificServiceArea(restaurants *[]model.Restaurant, area int) []primitive.ObjectID {
	set := make(map[primitive.ObjectID]bool)

	for _, res := range *restaurants {
		if isInServiceAreas(&res, area) {
			set[res.ID] = true
		}
	}

	return createSliceOfIDs(set)
}
