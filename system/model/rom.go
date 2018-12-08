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
	Machine []Machine		`bson:"machine"`
	Admin   []string 		`bson:"admin"`
	Acl 	bool			`bson:"acl"`
}

type Machine struct {
	Ip  string	`bson:"ip"`
	Mac string	`bson:"mac"`
	Des string	`bson:"bes"`
	Acl bool 	`bson:"acl"`
}
