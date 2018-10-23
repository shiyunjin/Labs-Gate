package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Machine struct {
	Id  int32
	Ip  string
	Mac string
	Des string
}

type Rom struct {
	Id      bson.ObjectId
	Name    string
	Code    string
	Vlan    string
	Device  string
	Machine []Machine
	Admin   []bson.ObjectId
}
