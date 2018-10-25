package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionUser = "user"
)

type User struct {
	Id         bson.ObjectId	`bson:"_id,omitempty"`
	Name       string
	Username   string
	Hash       string
	Salt       string
	Admin      bool
	Superadmin bool
	Permission string
	Rom        []string
	Createtime time.Time
	Updatetime time.Time
}
