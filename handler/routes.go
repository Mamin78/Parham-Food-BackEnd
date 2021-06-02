package handler

import (
	"github.com/labstack/echo/v4"
)

const (
	signUp = "/signup"
	login  = "/login"

	manager = "/manager"
	user    = "/user"

	timeline  = "/home"
	suggest   = "/suggestions"
	search    = "/search"
	userPath  = "/user"
	profiles  = "/profiles"
	tweets    = "/tweets"
	usernameQ = "/:username"
	follow    = usernameQ + "/follow"
	media     = "/media"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.GET("/", h.BaseRouter)

	g.POST("/restaurant", h.CreateRestaurant)
	g.POST(login, h.Login)
	g.POST("/temp", h.EditRestaurantInfo)
	//jwtMiddleware := middleware.JWT(utils.JWTSecret)
	//globalMiddleware := middleware.Global(utils.JWTSecret)
	//g.POST(signUp, h.SignUp)
	//g.POST(login, h.Login)
}
