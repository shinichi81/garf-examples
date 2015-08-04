package gapi

import (
	"github.com/weSolution/gapi-example/bundles/user"
	"github.com/weSolution/gapi/registry"
	"github.com/weSolution/gapi/server/echo"
)

func Registry() registry.Handler {

	// Creating new registry
	r := registry.New(echo.New())

	// Registering bundles
	r.Register(user.New())

	return r
}
