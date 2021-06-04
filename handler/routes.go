package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/router/middleware"
	"github.com/Mamin78/Parham-Food-BackEnd/utils"
	"github.com/labstack/echo/v4"
)

const (
	manager        = "/manager"
	restaurant     = "/restaurant"
	signUp         = "/signup"
	login          = "/login"
	edit           = "/edit"
	restaurantName = "/:restaurant_name"
	foodID         = "/:food_id"

	food  = "/food"
	order = "/order"
	foods = "/foods"
	user  = "/user"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.GET("/", h.BaseRouter)

	//sign up and login part!
	g.POST(manager+signUp, h.CreateRestaurantManager)
	g.POST(manager+login, h.ManagerLogin)

	g.POST(user+signUp, h.userSignUp)
	g.POST(user+login, h.UserLogin)

	//other parts
	managerGroup := g.Group(manager, middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				if c.Request().Method == "GET" {
					return true
				}
				//else if c.Path() == "/api"+manager+signUp || c.Path() == "/api"+manager+login {
				//	return true
				//}
				return false
			},
			SigningKey: utils.JWTSecret,
		},
	))
	managerGroup.POST("/create", h.CreateRestaurant)
	managerGroup.PUT("/update", h.EditRestaurantInfo)

	foodManager := managerGroup.Group(food)
	foodManager.POST("/add", h.CreateFood)
	foodManager.DELETE("/delete"+foodID, h.DeleteFood)
	foodManager.PUT("/disable"+foodID, h.DisableFood)
	foodManager.PUT("/enable"+foodID, h.EnableFood)
	//foodManager.PUT("/update"+foodID, h.CreateFood)

	orderManager := managerGroup.Group(order)
	orderManager.GET("/list", h.GetRestaurantOrders)

	orderStatusManager := orderManager.Group("/status")
	orderStatusManager.POST("/confirm"+"/:order_id", h.ConfirmOrderByRestaurantManager)

	//res := g.Group(restaurant, middleware.JWTWithConfig(
	//	middleware.JWTConfig{
	//		Skipper: func(c echo.Context) bool {
	//			if c.Request().Method == "GET" {
	//				return true
	//			}
	//			return false
	//		},
	//		SigningKey: utils.JWTSecret,
	//	},
	//))
	//res.GET(restaurantName, h.GetRestaurantInfo)
	//res.POST(restaurantName+food, h.CreateFood)
	//res.GET(restaurantName+foods, h.GetAllFoodsOfRestaurant)

	//jwtMiddleware := middleware.JWT(utils.JWTSecret)
	//userJWTMiddleware := middleware.USER(utils.JWTSecret)
	managerGroup.PUT(edit, h.EditRestaurantInfo)

	userGroup := g.Group(user, middleware.USERJWTFromHeader(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				//if c.Request().Method == "GET" {
				//	return true
				//}
				//else if c.Path() == "/api"+manager+signUp || c.Path() == "/api"+manager+login {
				//	return true
				//}
				return false
			},
			SigningKey: utils.JWTSecret,
		},
	))

	userGroup.PUT("/update", h.UpdateUserInfo)
	userGroup.GET("/info", h.GetUserInfo)
	userGroup.POST("/order", h.CreateOrder)
}
