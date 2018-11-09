package model

import "gopkg.in/mgo.v2/bson"

const (
	CollectionDevice = "device"
)

type Device struct {
	Id       bson.ObjectId	`bson:"_id"`
	Name     string
	Code     string
	Vlan     []string
	Invalid  []string
	Username string
	Password string
	Ip       string
}
