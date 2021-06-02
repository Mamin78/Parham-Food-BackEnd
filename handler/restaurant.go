package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2"
	"myapp/model"
	"net/http"
	"time"
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

	//here we should check the password!
	if res.Password != manager.Password {
		return c.JSON(http.StatusUnauthorized, "password is incorrect.")
	}

	////-----
	//// JWT
	////-----
	//
	//// Create token
	//token := jwt.New(jwt.SigningMethodHS256)
	//
	//// Set claims
	//claims := token.Claims.(jwt.MapClaims)
	//claims["id"] = manager.ID
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	//
	//// Generate encoded token and send it as response
	//manager.Token, err = token.SignedString([]byte(Key))
	//if err != nil {
	//	return err
	//}
	//
	//manager.Password = "" // Don't send password
	//return c.JSON(http.StatusOK, manager)

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	fmt.Println(res)
	fmt.Println("start claims part")
	fmt.Println(claims)
	fmt.Println(claims["id"])
	fmt.Println("start claims part")
	// Generate encoded token and send it as response
	res.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	res.Password = "" // Don't send password
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) EditRestaurantInfo(c echo.Context) (err error) {
	fmt.Println(c)
	managerEmail := restaurantManagerEmailFromToken(c)
	//id := c.Param("id")

	fmt.Println(managerEmail)
	// Add a follower to user

	return c.JSON(http.StatusOK, managerEmail)
}

func restaurantManagerEmailFromToken(c echo.Context) string {
	fmt.Println()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func stringFieldFromToken(c echo.Context, field string) string {
	field, ok := c.Get(field).(string)
	if !ok {
		return ""
	}
	return field
}
