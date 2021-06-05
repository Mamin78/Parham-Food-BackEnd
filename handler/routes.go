package handler

import (
	"github.com/Mamin78/Parham-Food-BackEnd/router/middleware"
	"github.com/Mamin78/Parham-Food-BackEnd/utils"
	"github.com/labstack/echo/v4"
)

const (
	URLManager        = "/manager"
	URLRestaurant     = "/restaurant"
	URLSignUp         = "/signup"
	URLLogin          = "/login"
	URLEdit           = "/edit"
	URLRestaurantName = "/:restaurant_name"
	URLFoodID         = "/:food_id"

	URLFood  = "/food"
	URLOrder = "/order"
	URLUser  = "/user"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {

	g.GET("/", h.BaseRouter)

	//sign up and login part!
	g.POST(URLManager+URLSignUp, h.CreateRestaurantManager)
	g.POST(URLManager+URLLogin, h.ManagerLogin)

	g.POST(URLUser+URLSignUp, h.userSignUp)
	g.POST(URLUser+URLLogin, h.UserLogin)

	//other parts
	managerGroup := g.Group(URLManager, middleware.JWTWithConfig(
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

	foodManager := managerGroup.Group(URLFood)
	foodManager.POST("/add", h.CreateFood)
	foodManager.DELETE("/delete"+URLFoodID, h.DeleteFood)
	foodManager.PUT("/disable"+URLFoodID, h.DisableFood)
	foodManager.PUT("/enable"+URLFoodID, h.EnableFood)
	//foodManager.PUT("/update"+foodID, h.CreateFood)

	orderManager := managerGroup.Group(URLOrder)
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
	managerGroup.PUT(URLEdit, h.EditRestaurantInfo)

	userGroup := g.Group(URLUser, middleware.USERJWTFromHeader(
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
	userComment := userGroup.Group("/comment")
	userComment.POST("/add", h.AddUserCommentToFood)

	foodGroup := g.Group(URLFood, middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				if c.Request().Method == "GET" {
					return true
				}
				return false
			},
			SigningKey: utils.JWTSecret,
		},
	))

	foodGroup.GET("/info"+"/:food_id", h.GetFoodInformation)
	foodGroup.GET("/comment"+"/:food_id", h.GetFoodComments)

}
