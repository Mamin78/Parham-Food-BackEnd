package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

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
	pass, err := PasswordToHash(manager.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "This user existed.", false))
	}
	res.Password = pass

	err = h.restaurantStore.CreateRestaurant(res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	return c.JSON(http.StatusCreated, model.NewResponse(newManagerResponse(manager), "", true))
}

func (h *Handler) CreateRestaurant(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Manager", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	if err := c.Bind(&res); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	res.Email = managerEmail
	res.BaseFoodTime = res.BaseFoodTime * time.Minute
	err = h.restaurantStore.UpdateInformation(managerEmail, res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
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

	//fmt.Println(res)
	//here we should check the password!
	//here we should check the password!
	if !PasswordsAreSame(res.Password, manager.Password) {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "password is incorrect.", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(newManagerResponse(manager), "", true))
}

func (h *Handler) EditRestaurantInfo(c echo.Context) (err error) {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Manager", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	newRes := model.NewRestaurant(res)
	if err := c.Bind(&newRes); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	err = h.restaurantStore.UpdateInformation(managerEmail, newRes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	newRes.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(newRes, "", true))
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

	//temp, err := h.foodsStore.GetAllFoodsOfRestaurantByID(res.ID)
	//if err != nil {
	//	if err == mgo.ErrNotFound {
	//		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid Restaurant!", false))
	//	}
	//	return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	//}
	//fmt.Println(temp)
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
	return c.JSON(http.StatusCreated, model.NewResponse(foods, "", true))
}

func stringFieldFromToken(c echo.Context, field string) string {
	field, ok := c.Get(field).(string)
	if !ok {
		return ""
	}
	return field
}
