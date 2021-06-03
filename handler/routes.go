package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"myapp/router/middleware"
	"myapp/utils"
)

const (
	manager        = "/manager"
	restaurant     = "/restaurant"
	signUp         = "/signup"
	login          = "/login"
	edit           = "/edit"
	restaurantName = "/:restaurant_name"

	food  = "/food"
	foods = "/foods"
	user  = "/user"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.GET("/", h.BaseRouter)

	g.POST(signUp+manager, h.CreateRestaurantManager)
	g.POST(login+manager, h.ManagerLogin)

	g.POST(signUp+user, h.CreateRestaurantManager)
	g.POST(login+user, h.ManagerLogin)

	res := g.Group(restaurant, middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				fmt.Println("in the keep function")
				if c.Request().Method == "GET" {
					return true
				}
				return false
			},
			SigningKey: utils.JWTSecret,
		},
	))
	res.GET(restaurantName, h.GetRestaurantInfo)
	res.POST(restaurantName+food, h.CreateFood)
	res.GET(restaurantName+foods, h.GetAllFoodsOfRestaurant)


	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	userJWTMiddleware := middleware.USER(utils.JWTSecret)

	manager := g.Group(manager, jwtMiddleware)
	manager.POST(restaurant, h.CreateRestaurant)
	manager.PUT(edit, h.EditRestaurantInfo)

	users := g.Group("/users", userJWTMiddleware)
	users.GET("", h.EditUser)
}
