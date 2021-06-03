package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"myapp/model"
	"net/http"
)

func (h *Handler) CreateRestaurant(c echo.Context) error {
	res := new(model.Restaurant)
	res.ID = primitive.NewObjectID()
	if err := c.Bind(&res); err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request1")
	}

	// Validate
	if res.Email == "" || res.Password == "" || len(res.Password) < 8 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	err := h.restaurantStore.CreateRestaurant(res)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request2")
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) Login(c echo.Context) error {
	manager := new(model.Restaurant)
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
	return c.JSON(http.StatusOK, newRestaurantResponse(res))
}

func (h *Handler) EditRestaurantInfo(c echo.Context) (err error) {
	//fmt.Println(c)
	//managerEmail := restaurantManagerEmailFromToken(c)
	//managerEmail1 := restaurantManagerEmailFromToken1(c)
	////id := c.Param("id")
	//
	//fmt.Println(managerEmail)
	//fmt.Println(managerEmail1)
	//fmt.Println("arman" + stringFieldFromToken(c, "email"))
	//fmt.Println("tempppp" + userPhoneFromToken(c))
	//// Add a follower to user

	fmt.Println(c)
	return c.JSON(http.StatusOK, stringFieldFromToken(c, "email"))
}

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
	//fmt.Println(c)
	//managerEmail := restaurantManagerEmailFromToken(c)
	//managerEmail1 := restaurantManagerEmailFromToken1(c)
	////id := c.Param("id")
	//
	//fmt.Println(managerEmail)
	//fmt.Println(managerEmail1)
	//fmt.Println("arman" + stringFieldFromToken(c, "email"))
	//fmt.Println("tempppp" + userPhoneFromToken(c))
	//// Add a follower to user

	fmt.Println(c)
	return c.JSON(http.StatusOK, stringFieldFromToken(c, "phone"))
}

//func restaurantManagerEmailFromToken(c echo.Context) string {
//	fmt.Println()
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	return claims["id"].(string)
//}
//
//func restaurantManagerEmailFromToken1(c echo.Context) string {
//	fmt.Println()
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//	return claims["email"].(string)
//}
//
//func userPhoneFromToken(c echo.Context) string {
//	fmt.Println()
//	user := c.Get("user").(*jwt.Token)
//	claims := user.Claims.(jwt.MapClaims)
//
//	if _, ok := claims["phone"]; !ok {
//		return ""
//	}
//	return claims["phone"].(string)
//}
//
//func userIDFromToken(c echo.Context) uint {
//	id, ok := c.Get("user").(uint)
//	if !ok {
//		return 0
//	}
//	return id
//}

func stringFieldFromToken(c echo.Context, field string) string {
	field, ok := c.Get(field).(string)
	if !ok {
		return ""
	}
	return field
}
