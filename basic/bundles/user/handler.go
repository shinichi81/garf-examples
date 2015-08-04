package user

import (
	"github.com/backenderia/garf/server"
	"gopkg.in/mgo.v2/bson"
)

func (u *bundle) prepare(c server.Context) User {
	id := u.server.Param(c, "id")
	name := u.server.Form(c, "name")

	q := User{
		ID:   bson.ObjectIdHex(id),
		Name: name,
	}

	return q
}

func (u *bundle) handler(f ModelHandler) server.HttpHandler {
	return func(c server.Context) error {
		query := u.prepare(c)

		result, err := f(query)
		if err != nil {
			return u.server.Error(c, 500, err)
		}

		return u.server.JSON(c, result)
	}
}
