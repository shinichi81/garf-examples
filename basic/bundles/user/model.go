package user

import "gopkg.in/mgo.v2/bson"

// ModelHandler represents all model functions on this package
type ModelHandler func(User) ([]User, error)

// List User
func List(q User) (res []User, err error) {
	var result []User
	user, conn := Db()
	defer conn.Close()
	err = user.Find(bson.M{}).All(&result)
	return result, err
}

// Create User
func Create(q User) (res []User, err error) {
	user, conn := Db()
	defer conn.Close()
	err = user.Insert(&User{
		Name: q.Name,
	})
	return
}

// Read User
func Read(q User) ([]User, error) {
	var result User
	user, conn := Db()
	defer conn.Close()
	err := user.Find(bson.M{
		"_id": q.ID,
	}).One(&result)
	return []User{result}, err
}

// Update User
func Update(q User) (res []User, err error) {
	user, conn := Db()
	defer conn.Close()
	err = user.Update(
		bson.M{
			"_id": q.ID,
		},
		bson.M{
			"$set": bson.M{
				"Name": q.Name,
			},
		},
	)
	return
}

// Delete User
func Delete(q User) (res []User, err error) {
	user, conn := Db()
	defer conn.Close()
	err = user.RemoveId(q.ID)
	return
}
