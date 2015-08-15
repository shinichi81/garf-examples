package main

import (
	"github.com/backenderia/garf-contrib/jwt-auth"
	"github.com/backenderia/garf-examples/basic/bundles/user"
	"github.com/backenderia/garf-examples/basic/garf"
	"github.com/dgrijalva/jwt-go"
	mw "github.com/labstack/echo/middleware"
)

const (
	signingKey = "mysecretkey"
)

func main() {

	r := garf.Registry()
	r.Set("URI", "mongodb://localhost:27017")

	s := r.Server()

	s.Use(mw.Logger())
	s.Use(mw.Recover())

	// JWT Authentication middleware
	s.Use(jwtauth.New(func(token *jwt.Token) (sign interface{}, err error) {
		id := token.Claims["id"]
		secret := token.Claims["secret"]

		if id, ok := id.(string); ok {
			if secret, ok := secret.(string); ok {
				success := false
				success, err = user.Auth(user.User{
					ID:     id,
					Secret: []byte(secret),
				})

				if success {
					sign = []byte(signingKey)
					return
				}
			}
		}

		return
	}))

	r.Configure()

	// ** DEMO INSTRUCTIONS: **
	//
	// Uncomment this part below to create the demo user on your mongo server
	//
	// user.Create(user.User{
	// 	ID:     "user",
	// 	Secret: []byte("123"),
	// })
	//
	// Then use this Authentication token for requesting:
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6InVzZXIiLCJzZWNyZXQiOiIxMjMifQ.Affo3QE5iLDgXfllo_o47uEpUJp-qRAWexVZ8ZdfRZQ
	//
	// Note: That token works only if SingingKey="mysecretkey"

	s.Run(":3000")
}
