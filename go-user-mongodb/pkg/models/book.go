package models

import (
	"fmt"

	"github.com/go-user-mongodb/pkg/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Session

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetUser(Id string) (*User, *mgo.Session) {
	var getUser User
	oid := bson.ObjectIdHex(Id)
	if err := db.DB("mongo-golang").C("users").FindId(oid).One(&getUser); err != nil {
		return &User{}, db
	}
	return &getUser, db
}

func (b *User) CreateUser() *User {
	b.Id = bson.NewObjectId()
	fmt.Println("------------------------")
	fmt.Print(db)
	db.DB("mongo-golang").C("users").Insert(b)
	return b
}
