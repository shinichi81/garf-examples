package garf

import (
	"github.com/backenderia/garf-example/bundles/user"
	"github.com/backenderia/garf/registry"
	"github.com/backenderia/garf/server/echo"
)

func Registry() registry.Handler {

	// Creating new registry
	r := registry.New(echo.New())

	// Registering bundles
	r.Register(user.New())

	return r
}
