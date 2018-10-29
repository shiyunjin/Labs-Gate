package model

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionRom = "rom"
)

type Roms struct {
	Id      bson.ObjectId	`bson:"_id,omitempty"`
	Name 	string
	Rom 	[]Rom
}

type Rom struct {
	Name    string
	Code    string
	Vlan    string
	Device  string
	Machine []Machine
	Admin   []bson.ObjectId
}

type Machine struct {
	Ip  string
	Mac string
	Des string
}
