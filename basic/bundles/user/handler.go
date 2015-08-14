package user

import "github.com/backenderia/garf/server"

func (u *bundle) prepare(c server.Context) User {

	id := u.server.Param(c, "id")
	secret := u.server.Form(c, "secret")

	q := User{
		ID:     id,
		Secret: []byte(secret),
	}

	return q
}

func (u *bundle) handler(f ModelHandler) server.HandlerFunc {
	return func(c server.Context) error {
		query := u.prepare(c)

		result, err := f(query)
		if err != nil {
			return u.server.Error(c, 500, err)
		}

		return u.server.JSON(c, result)
	}
}
