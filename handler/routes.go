package handler

import (
	"github.com/labstack/echo/v4"
	"myapp/router/middleware"
	"myapp/utils"
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

	g.POST("/signup", h.userSignUp)
	g.POST("/userlogin", h.UserLogin)
	//g.POST("/temp", h.EditRestaurantInfo)
	//jwtMiddleware := middleware.JWT(utils.JWTSecret)
	//globalMiddleware := middleware.Global(utils.JWTSecret)
	//g.POST(signUp, h.SignUp)
	//g.POST(login, h.Login)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	userJWTMiddleware := middleware.USER(utils.JWTSecret)
	//guestUsers := g.Group("/users")
	//guestUsers.POST("", h.CreateRestaurant)
	//guestUsers.POST("/login", h.Login)

	user := g.Group("/user", jwtMiddleware)
	user.GET("", h.EditRestaurantInfo)

	users := g.Group("/users", userJWTMiddleware)
	users.GET("", h.EditUser)
	//user.PUT("", h.UpdateUser)
	//
	//profiles := g.Group("/profiles", jwtMiddleware)
	//profiles.GET("/:username", h.GetProfile)
	//profiles.POST("/:username/follow", h.Follow)
	//profiles.DELETE("/:username/follow", h.Unfollow)
	//
	//articles := g.Group("/articles", middleware.JWTWithConfig(
	//	middleware.JWTConfig{
	//		Skipper: func(c echo.Context) bool {
	//			if c.Request().Method == "GET" && c.Path() != "/api/articles/feed" {
	//				return true
	//			}
	//			return false
	//		},
	//		SigningKey: utils.JWTSecret,
	//	},
	//))
	//articles.POST("", h.CreateArticle)
	//articles.GET("/feed", h.Feed)
	//articles.PUT("/:slug", h.UpdateArticle)
	//articles.DELETE("/:slug", h.DeleteArticle)
	//articles.POST("/:slug/comments", h.AddComment)
	//articles.DELETE("/:slug/comments/:id", h.DeleteComment)
	//articles.POST("/:slug/favorite", h.Favorite)
	//articles.DELETE("/:slug/favorite", h.Unfavorite)
	//articles.GET("", h.Articles)
	//articles.GET("/:slug", h.GetArticle)
	//articles.GET("/:slug/comments", h.GetComments)
	//
	//tags := g.Group("/tags")
	//tags.GET("", h.Tags)
}
