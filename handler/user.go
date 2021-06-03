package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"myapp/model"
	"net/http"
)

func (h *Handler) userSignUp(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request1")
	}

	// Validate
	if user.PhoneNumber == "" || user.Password == "" || len(user.Password) < 8 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid phone or password"}
	}

	user.ID = primitive.NewObjectID()
	err := h.userStore.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request2")
	}
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) UserLogin(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	res, err := h.userStore.GetUserByPhone(user.PhoneNumber)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid user")
		}
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	fmt.Println(res)
	//here we should check the password!
	if res.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, "password is incorrect.")
	}
	return c.JSON(http.StatusOK, newUserResponse(res))
}

func (h *Handler) EditUser(c echo.Context) (err error) {
	fmt.Println(c)
	return c.JSON(http.StatusOK, stringFieldFromToken(c, "phone"))
}
