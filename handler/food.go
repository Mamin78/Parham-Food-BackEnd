package handler

import (
	"fmt"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"net/http"
)

func (h *Handler) CreateFood(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	res.Password = ""

	food := new(model.Food)
	food.ID = primitive.NewObjectID()
	food.RestaurantID = res.ID
	food.CanBeOrdered = true

	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	err = h.foodsStore.CreateFood(food)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.restaurantStore.AddFoodToRestaurant(res.Name, food, res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(addRestaurantNameToFood(h, food), "", true))
}

func (h *Handler) DisableFood(c echo.Context) error {
	foodID := c.Param("food_id")

	err := h.foodsStore.DisableFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "food not found", false))
	}

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "food not found", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(addRestaurantNameToFood(h, food), "", true))
}

func (h *Handler) EnableFood(c echo.Context) error {
	foodID := c.Param("food_id")

	err := h.foodsStore.EnableFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(addRestaurantNameToFood(h, food), "", true))
}

func (h *Handler) DeleteFood(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	res.Password = ""

	foodID := c.Param("food_id")

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.foodsStore.DeleteFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	err = h.restaurantStore.DeleteFoodFromRestaurant(foodID, res)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid restaurant", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(food, "", true))
}

func (h *Handler) GetFoodInformation(c echo.Context) error {
	foodID := c.Param("food_id")

	food, err := h.foodsStore.GetFoodByID(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(addRestaurantNameToFood(h, food), "", true))
}

func (h *Handler) GetFoodComments(c echo.Context) error {
	foodID := c.Param("food_id")

	fmt.Println(foodID)
	comments, err := h.commentsStore.GetAllFoodComments(foodID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid food", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(addUserNameToComment(h, comments), "", true))
}

func addUserNameToComment(h *Handler, comments *[]model.Comment) []model.CommentAsResponse {
	var result []model.CommentAsResponse
	for _, v := range *comments {
		commentWithUserName := model.CreateRespCommentFromComment(v)

		user, err := h.userStore.GetUserByPrimitiveID(v.UserID)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			return nil
		}

		commentWithUserName.UserName = user.Name
		result = append(result, *commentWithUserName)
	}
	return result
}

func addRestaurantNameToFood(h *Handler, food *model.Food) *model.FoodAsResponse {
	foodWithResName := model.CreateRepFoodWithResName(*food)

	res, err := h.restaurantStore.GetRestaurantByPrimitiveTypeId(food.RestaurantID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		return nil
	}

	foodWithResName.RestaurantName = res.Name

	return foodWithResName
}
