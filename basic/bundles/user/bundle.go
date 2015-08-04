package user

import (
	"log"

	"github.com/backenderia/garf/registry"
	"github.com/backenderia/garf/server"
	"github.com/backenderia/garf-contrib/mongodb"
	"gopkg.in/mgo.v2"
)

const (
	basePath = "/user/"
)

type bundle struct {
	server server.Handler
}

// New instance from this bundle
func New() registry.Bundle {
	return &bundle{}
}

// Db retrieves User's collection and session
var Db func() (*mgo.Collection, *mgo.Session)

// Init bundle
func (u *bundle) Init(c map[string]interface{}) {
	Db = mongodb.Configure(c["uri"].(string), "User")
	log.Println("`User` route configured...")
}

// Register bundle's routes to server
func (u *bundle) Register(r server.Handler) {
	u.server = r

	route := r.Group(basePath)
	route.Get("", u.handler(List))
	route.Post("", u.handler(Create))
	route.Get(":id", u.handler(Read))
	route.Post(":id", u.handler(Update))
	route.Del(":id", u.handler(Delete))
}
