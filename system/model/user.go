package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId
	Name       string
	Username   string
	Hash       string
	Salt       string
	Admin      bool
	Superadmin bool
	Permission []string
	Rom        []string
	Createtime string
	Updatetime string
}