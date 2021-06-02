package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

//func New() *echo.Echo {
//	e := echo.New()
//	e.Logger.SetLevel(log.DEBUG)
//
//	e.Logger.SetLevel(log.ERROR)
//	e.Use(middleware.Logger())
//	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
//		SigningKey: []byte(handler.Key),
//		Skipper: func(c echo.Context) bool {
//			// Skip authentication for signup and login requests
//			if c.Path() == "/login" || c.Path() == "/signup" || c.Path() == "/restaurant" {
//				return true
//			}
//			return false
//		},
//	}))
//
//	e.Pre(middleware.RemoveTrailingSlash())
//	e.Use(middleware.Logger())
//	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//		AllowOrigins: []string{"*"},
//		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
//		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
//	}))
//	e.Validator = NewValidator()
//	return e
//}

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Validator = NewValidator()
	return e
}