package handler

import (
	"github.com/labstack/echo/v4"
	"myapp/router/middleware"
	"myapp/utils"
)

const (
	manager      = "/manager"
	restaurant   = "/restaurant"
	signUp       = "/signup"
	login        = "/login"
	edit         = "/edit"
	restaurantId = "/:restaurant_id"

	user = "/user"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.GET("/", h.BaseRouter)

	g.POST(signUp+manager, h.CreateRestaurantManager)
	g.POST(login+manager, h.ManagerLogin)

	g.POST(signUp+user, h.CreateRestaurantManager)
	g.POST(login+user, h.ManagerLogin)

	res := g.Group(restaurant)
	res.GET(restaurantId, h.GetRestaurantInfo)

	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	//userJWTMiddleware := middleware.USER(utils.JWTSecret)

	manager := g.Group(manager, jwtMiddleware)
	manager.POST(restaurant, h.CreateRestaurant)
	manager.PUT(edit, h.EditRestaurantInfo)

	//users := g.Group("/users", userJWTMiddleware)
	//users.GET("", h.EditUser)
}
