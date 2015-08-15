package user

import (
	"bytes"
	"crypto/sha1"

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
		ID:     q.ID,
		Secret: createHash(q.Secret),
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
				"Secret": createHash(q.Secret),
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
