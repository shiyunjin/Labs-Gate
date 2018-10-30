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
	Rom		model.Rom
}

func RomGet(db *mgo.Database, code string) (result RomResponse,err error) {

	err = db.C(model.CollectionRom).Pipe([]bson.M{
		{"$unwind": "$rom"},
		{"$match": bson.M{"rom.code": code}},
		{"$project": bson.M{"rom": 1,"_id": 0}},
	}).One(&result)

	return result, err
}

func OpenRom(c *gin.Context) {
	code := c.Param("code")

	// TODO: add user permission auth

	db := c.MustGet("db").(*mgo.Database)
	result, err := RomGet(db, code)

	// Log Rom data for test
	fmt.Println(result.Rom.Machine)

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

	db := c.MustGet("db").(*mgo.Database)
	result, err := RomGet(db, code)

	// Log Rom data for test
	fmt.Println(result.Rom.Machine)

	// TODO: each machine to connect server with telnet for close network

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
	})
}