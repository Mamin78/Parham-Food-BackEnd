package handler

import (
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

	userInformation, err := h.userStore.GetUserByPhone(user.PhoneNumber)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid user")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	//here we should check the password!
	if userInformation.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, "password is incorrect.")
	}
	return c.JSON(http.StatusOK, newUserResponse(userInformation))
}

func (h *Handler) UpdateUserInfo(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid user")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	newUser := model.NewUser(userInformation)

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	err = h.userStore.UpdateUserInfoByPhone(userPhone, newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request3")
	}
	newUser.Password = ""
	return c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) GetUserInfo(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusUnauthorized, "invalid user")
		}
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	userInformation.Password = ""
	return c.JSON(http.StatusCreated, userInformation)
}
