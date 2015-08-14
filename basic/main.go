package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/backenderia/garf-examples/basic/bundles/user"
	"github.com/backenderia/garf-examples/basic/garf"
	"github.com/backenderia/garf/server"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {

	r := garf.Registry()
	r.Set("URI", "mongodb://localhost:27017")

	s := r.Server()

	s.Use(mw.Logger())
	s.Use(mw.Recover())

	s.Use(JWTAuth(func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Claims["id"].(string); !ok {
			return nil, errors.New("missing id")
		}

		u, err := user.Read(user.User{
			ID: token.Claims["id"].(string),
		})

		if len(u) == 0 || err != nil {
			return nil, err
		}

		return u[0].Secret, nil
	}))

	r.Configure()
	s.Run(":3000")
}

type SecretFunc func(*jwt.Token) (interface{}, error)

func JWTAuth(a SecretFunc) server.HandlerFunc {
	return func(c server.Context) error {
		if (c.Request().Header.Get(echo.Upgrade)) == echo.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get("Authorization")
		l := len("Bearer")
		he := echo.NewHTTPError(http.StatusUnauthorized)

		if len(auth) > l+1 && auth[:l] == "Bearer" {
			t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return a(token)
			})
			if err == nil && t.Valid {
				c.Set("claims", t.Claims)
				return nil
			}
		}
		return he
	}
}
