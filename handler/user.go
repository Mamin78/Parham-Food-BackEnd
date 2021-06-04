package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"net/http"
)

func (h *Handler) userSignUp(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	// Validate
	if user.PhoneNumber == "" || user.Password == "" || len(user.Password) < 8 {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid phone or password", false))
	}

	user.ID = primitive.NewObjectID()
	err := h.userStore.CreateUser(user)
	user.Credit = 1000000
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	user.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(user, "", true))
}

func (h *Handler) UserLogin(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	userInformation, err := h.userStore.GetUserByPhone(user.PhoneNumber)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	//here we should check the password!
	if userInformation.Password != user.Password {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "password is incorrect.", false))
	}
	return c.JSON(http.StatusCreated, model.NewResponse(newUserResponse(userInformation), "", true))
}

func (h *Handler) UpdateUserInfo(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	newUser := model.NewUser(userInformation)

	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	err = h.userStore.UpdateUserInfoByPhone(userPhone, newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}
	newUser.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(newUser, "", true))
}

func (h *Handler) GetUserInfo(c echo.Context) (err error) {
	userPhone := stringFieldFromToken(c, "phone")

	userInformation, err := h.userStore.GetUserByPhone(userPhone)
	if err != nil {
		if err == mgo.ErrNotFound {
			return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "invalid user", false))
		}
		return c.JSON(http.StatusBadRequest, model.NewResponse(nil, "bad request", false))
	}

	userInformation.Password = ""
	return c.JSON(http.StatusCreated, model.NewResponse(userInformation, "", true))
}
