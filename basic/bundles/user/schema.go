package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name,omitempty"`
}
