package config

import (
	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Session
)

func Connect() {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	db = s
}

func GetDB() *mgo.Session {
	return db
}
