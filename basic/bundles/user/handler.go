package user

import "github.com/backenderia/garf/server"

func (u *bundle) prepare(c server.Context) User {

	id := c.Param("id")
	secret := c.Form("secret")

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
			return c.Error(500, err)
		}

		return c.JSON(200, result)
	}
}
