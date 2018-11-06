package db

import (
	"fmt"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// Connect connects to mongodb
func Connect() {

	uri := "mongodb://" + config.Get("mongodb.host").(string) +
		":" + config.Get("mongodb.port").(string) +
		"/" + config.Get("mongodb.name").(string)

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}

	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}