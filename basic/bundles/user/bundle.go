package user

import (
	"log"

	"github.com/backenderia/garf-contrib/mongodb"
	"github.com/backenderia/garf/registry"
	"github.com/backenderia/garf/server"
	"gopkg.in/mgo.v2"
)

const (
	basePath = "/user"
)

type bundle struct{}

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
func (u *bundle) Register(r server.Support) {
	r.Get(basePath+"/user/", u.handler(List))
	r.Post(basePath+"/", u.handler(Create))
	r.Get(basePath+"/:id", u.handler(Read))
	r.Post(basePath+"/:id", u.handler(Update))
	r.Del(basePath+"/:id", u.handler(Delete))
}
