package handler

import (
	"github.com/labstack/echo/v4"
	"myapp/router/middleware"
	"myapp/utils"
)

const (
	manager    = "/manager"
	restaurant = "/restaurant"
	signUp     = "/signup"
	login      = "/login"
	edit       = "/edit"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.GET("/", h.BaseRouter)

	g.POST(manager, h.CreateRestaurantManager)
	g.POST(manager+restaurant, h.CreateRestaurant)
	g.POST(manager+login, h.CreateRestaurantManager)

	g.POST("/signup", h.userSignUp)
	g.POST("/userlogin", h.UserLogin)

	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	userJWTMiddleware := middleware.USER(utils.JWTSecret)

	manager := g.Group("", jwtMiddleware)
	manager.GET(edit, h.EditRestaurantInfo)

	user := g.Group("/user", jwtMiddleware)
	user.GET("", h.EditRestaurantInfo)

	users := g.Group("/users", userJWTMiddleware)
	users.GET("", h.EditUser)
}
