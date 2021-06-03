package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"myapp/model"
	"net/http"
)

func (h *Handler) userSignUp(c echo.Context) error {
	res := new(model.User)
	if err := c.Bind(&res); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request1")
	}

	// Validate
	if res.PhoneNumber == "" || res.Password == "" || len(res.Password) < 8 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid phone or password"}
	}

	err := h.userStore.CreateUser(res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request2")
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) UserLogin(c echo.Context) error {
	manager := new(model.User)
	if err := c.Bind(&manager); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request1user")
	}

	res, err := h.userStore.GetUserByPhone(manager.PhoneNumber)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid email or password")
		}
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, "Bad Request 2 user")
	}

	fmt.Println(res)
	//here we should check the password!
	if res.Password != manager.Password {
		return c.JSON(http.StatusUnauthorized, "password is incorrect.")
	}
	return c.JSON(http.StatusOK, newUserResponse(res))
}

func (h *Handler) EditUser(c echo.Context) (err error) {
	fmt.Println(c)
	return c.JSON(http.StatusOK, stringFieldFromToken(c, "phone"))
}
