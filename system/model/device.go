package model

import "gopkg.in/mgo.v2/bson"

type Device struct {
	Id       bson.ObjectId
	Name     string
	Code     string
	Vlan     []string
	Invalid  []string
	Username string
	Password string
	Ip       string
}
