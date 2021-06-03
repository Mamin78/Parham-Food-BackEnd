package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"myapp/utils"
	"net/http"
)

type (
	USERJWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	USERSkipper      func(c echo.Context) bool
	USERjwtExtractor func(echo.Context) (string, error)
)

var (
	USERErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
)

func USER(key interface{}) echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = key
	return USERJWTFromHeader(c)
}

func USERJWTFromHeader(config JWTConfig) echo.MiddlewareFunc {
	extractor := jwtFromHeader("Authorization", "Token")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractor(c)
			if auth == "" && config.Skipper != nil && config.Skipper(c) {
				return next(c)
			}
			if err != nil {
				return c.JSON(http.StatusForbidden, utils.NewError(USERErrJWTInvalid))
			}

			if auth == "" {
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return c.JSON(http.StatusUnauthorized, utils.NewError(errors.New("missing or malformed jwt")))
			}
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})
			if err != nil {
				return c.JSON(http.StatusForbidden, utils.NewError(USERErrJWTInvalid))
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				phone := claims["phone"]
				c.Set("phone", phone)
				return next(c)
			}
			return c.JSON(http.StatusForbidden, utils.NewError(USERErrJWTInvalid))
		}
	}
}