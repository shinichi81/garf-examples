package user

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"gopkg.in/mgo.v2/bson"
)

// ModelHandler represents all model functions on this package
type ModelHandler func(User) ([]User, error)

const saltSize = 16

// GenerateSalt creates a new salt for hashing the password
func GenerateSalt(secret []byte) []byte {
	buf := make([]byte, saltSize, saltSize+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)

	if err != nil {
		fmt.Printf("random read failed: %v", err)
		os.Exit(1)
	}

	hash := sha1.New()
	hash.Write(buf)
	hash.Write(secret)
	return hash.Sum(buf)
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
		Secret: GenerateSalt(q.Secret),
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
				"Secret": GenerateSalt(q.Secret),
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
