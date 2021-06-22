package handler

import (
	"fmt"
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

const form = "15:04"

func (h *Handler) CreateRestaurantManager(c echo.Context) error {
	manager := new(model.BaseManager)
	if err := c.Bind(&manager); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	if manager.Email == "" {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "phone number must not be empty", false))
	}

	if err := CheckPasswordLever(manager.Password); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, err.Error(), false))
	}

	res := new(model.Restaurant)
	res.ID = primitive.NewObjectID()
	res.Email = manager.Email
	res.Name = res.ID.String()
	pass, err := PasswordToHash(manager.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "This user existed.", false))
	}
	res.Password = pass

	err = h.restaurantStore.CreateRestaurant(res)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "This email is already registered", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(newManagerResponse(manager), "", true))
}

func (h *Handler) CreateRestaurant(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	var restaurant *model.BaseRestaurant
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Manager", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Manager", false))
	}
	if err := c.Bind(&restaurant); err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	model.NewRestaurantFromBaseRes(res, restaurant)
	res.StartWorkingHours, _ = time.Parse(form, restaurant.StartWorkingHours)
	res.EndWorkingHours, _ = time.Parse(form, restaurant.EndWorkingHours)

	fmt.Println(res.StartWorkingHours.Clock())
	//temp := res.StartWorkingHours.Clock()
	//temp.
	res.Email = managerEmail
	res.BaseFoodTime = res.BaseFoodTime * time.Minute
	err = h.restaurantStore.UpdateInformation(managerEmail, res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant", false))
	}
	res.Password = ""

	return c.JSON(http.StatusCreated, model.NewResponse(res, "", true))
}

func (h *Handler) ManagerLogin(c echo.Context) error {
	manager := new(model.BaseManager)
	if err := c.Bind(&manager); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(manager.Email)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid email or password", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	//here we should check the password!
	if !PasswordsAreSame(res.Password, manager.Password) {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "password is incorrect.", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(newManagerResponse(manager), "", true))
}

func (h *Handler) EditRestaurantInfo(c echo.Context) (err error) {
	managerEmail := stringFieldFromToken(c, "email")
	var restaurant *model.BaseRestaurant

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Manager", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	newRes := model.NewRestaurant(res)
	if err := c.Bind(&restaurant); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	model.NewRestaurantFromBaseRes(newRes, restaurant)
	res.StartWorkingHours, _ = time.Parse(form, restaurant.StartWorkingHours)
	res.EndWorkingHours, _ = time.Parse(form, restaurant.EndWorkingHours)

	err = h.restaurantStore.UpdateInformation(managerEmail, newRes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	newRes.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(newRes, "", true))
}

func (h *Handler) GetRestaurantInfoByToken(c echo.Context) (err error) {
	managerEmail := stringFieldFromToken(c, "email")

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
	}
	res.Password = ""

	return c.JSON(http.StatusCreated, model.NewResponse(res, "", true))
}

func (h *Handler) GetRestaurantInfo(c echo.Context) (err error) {
	resName := c.Param("restaurant_name")
	res, err := h.restaurantStore.GetRestaurantByName(resName)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	res.Password = ""

	return c.JSON(http.StatusCreated, model.NewResponse(res, "", true))
}

func (h *Handler) GetAllFoodsOfRestaurant(c echo.Context) (err error) {
	resName := c.Param("restaurant_name")
	res, err := h.restaurantStore.GetRestaurantByName(resName)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	res.Password = ""

	foods, err := h.foodsStore.GetAllFoodsOfRestaurantByID(res.ID)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(addRestaurantNameToFoods(h, &foods), "", true))
}

func stringFieldFromToken(c echo.Context, field string) string {
	field, ok := c.Get(field).(string)
	if !ok {
		return ""
	}
	return field
}

func addRestaurantNameToFoods(h *Handler, foods *[]model.Food) []model.FoodAsResponse {
	var result []model.FoodAsResponse
	for _, v := range *foods {
		foodWithResName := model.CreateRepFoodWithResName(v)

		res, err := h.restaurantStore.GetRestaurantByPrimitiveTypeId(v.RestaurantID)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			return nil
		}

		foodWithResName.RestaurantName = res.Name
		result = append(result, *foodWithResName)
	}
	return result
}
