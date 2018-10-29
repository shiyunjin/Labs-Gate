package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RomResponse struct {

}

func OpenRom(c *gin.Context) {
	code := c.Param("code")

	// TODO: add user permission auth

	result := bson.M{}
	db := c.MustGet("db").(*mgo.Database)
	err := db.C(model.CollectionRom).Pipe([]bson.M{
		{"$unwind": "$rom"},
		{"$match": bson.M{"rom.code": code}},
		{"$project": bson.M{"rom": 1,"_id": 0}},
	}).One(&result)

	//rom := result["rom"]

	// Log Rom data for test
	//fmt.Println()

	// TODO: each machine to connect server with telnet for open network

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
	})
}

func CloseRom(c *gin.Context) {
	code := c.Param("code")

	// TODO: add user permission auth

	result := model.Roms{}
	db := c.MustGet("db").(*mgo.Database)
	err := db.C(model.CollectionRom).Find(nil).Select(bson.M{
		"rom": bson.M{
			"$elemMatch": bson.M{
				"code": code,
			},
		},
	}).One(&result)

	rom := result.Rom[0]

	// Log Rom data for test
	fmt.Printf("%v", rom)

	// TODO: each machine to connect server with telnet for close network

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
	})
}