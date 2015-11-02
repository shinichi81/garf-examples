package user

import (
	"bytes"
	"crypto/sha1"

	"github.com/backenderia/garf-contrib/adapter"
	"gopkg.in/mgo.v2/bson"
)

// ModelHandler represents all model functions on this package
type ModelHandler func(User) ([]User, error)

const saltSize = 16

func createHash(secret []byte) []byte {
	result := make([]byte, saltSize)
	hash := sha1.New()
	hash.Write(secret)
	return hash.Sum(result)
}

// Auth given User
func Auth(u User) (bool, error) {
	user, err := Read(User{
		ID: u.ID,
	})

	if len(user) == 0 || err != nil {
		return false, err
	}

	secret := createHash(u.Secret)
	if !bytes.Equal(user[0].Secret, secret) {
		return false, nil
	}

	return true, nil
}

// List User
func List(q User) (res []User, err error) {
	var result []User

	task := UserStore.Read(adapter.M{}).Options(adapter.M{
		"All": true,
	})

	err = task.Exec(result)

	return result, err
}

// Create User
func Create(q User) (res []User, err error) {
	task := UserStore.Create(q.ID, adapter.M{
		"secret": createHash(q.Secret),
	})

	var u User
	err = task.Exec(&u)

	return []User{u}, err
}

// Read User
func Read(q User) (res []User, err error) {
	task := UserStore.Read(adapter.M{
		"_id": q.ID,
	})

	var u User
	err = task.Exec(&u)

	return []User{u}, err
}

// Update User
func Update(q User) (res []User, err error) {
	task := UserStore.Update(adapter.M{
		"_id": q.ID,
	}, adapter.M{
		"$set": bson.M{
			"Secret": createHash(q.Secret),
		},
	})

	err = task.Exec(&res)

	return
}

// Delete User
func Delete(q User) (res []User, err error) {
	task := UserStore.Delete(adapter.M{
		"_id": q.ID,
	})

	err = task.Exec(nil)

	return
}
