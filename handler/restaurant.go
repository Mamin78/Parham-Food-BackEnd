package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"myapp/model"
	"net/http"
)

func (h *Handler) CreateRestaurantManager(c echo.Context) error {
	manager := new(model.BaseManager)
	if err := c.Bind(&manager); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Validate
	if manager.Email == "" || manager.Password == "" || len(manager.Password) < 8 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	res := new(model.Restaurant)
	//res.ID = primitive.NewObjectID()
	res.Email = manager.Email
	res.Password = manager.Password

	err := h.restaurantStore.CreateRestaurant(res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusCreated, newManagerResponse(manager))
}

func (h *Handler) CreateRestaurant(c echo.Context) error {
	managerEmail := stringFieldFromToken(c, "email")
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)
	if err := c.Bind(&res); err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Manager")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request1")
	}

	res.Email = managerEmail
	err = h.restaurantStore.UpdateInformation(managerEmail, res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request2")
	}
	res.Password = ""
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) ManagerLogin(c echo.Context) error {
	manager := new(model.BaseManager)
	if err := c.Bind(&manager); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	res, err := h.restaurantStore.GetRestaurantByManagerEmail(manager.Email)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid email or password")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	fmt.Println(res)
	//here we should check the password!
	if res.Password != manager.Password {
		return c.JSON(http.StatusUnauthorized, "password is incorrect.")
	}
	return c.JSON(http.StatusOK, newManagerResponse(manager))
}

func (h *Handler) EditRestaurantInfo(c echo.Context) (err error) {
	managerEmail := stringFieldFromToken(c, "email")
	fmt.Println("emailllll" + managerEmail)
	res, err := h.restaurantStore.GetRestaurantByManagerEmail(managerEmail)

	fmt.Println("ressssss " )
	fmt.Println( res)
	newRes := model.NewRestaurant(res)
	fmt.Println("newressssss " )
	fmt.Println( newRes)
	if err := c.Bind(&newRes); err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid Manager")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request1")
	}
	fmt.Println("newwwwwwwwwwressssss " )
	fmt.Println( newRes)
	err = h.restaurantStore.UpdateInformation(managerEmail, newRes)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request2")
	}
	newRes.Password = ""
	return c.JSON(http.StatusCreated, newRes)
}

func stringFieldFromToken(c echo.Context, field string) string {
	field, ok := c.Get(field).(string)
	if !ok {
		return ""
	}
	return field
}
