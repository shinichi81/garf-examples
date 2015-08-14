package user

type User struct {
	ID     string `bson:"_id,omitempty"`
	Secret []byte `bson:"secret,omitempty" json:"-"`
}
