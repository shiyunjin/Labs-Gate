package network

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"github.com/shiyunjin/Labs-Gate/system/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RomResponse struct {
	Rom		model.Rom
}

type MachineResponse struct {
	Rom 	struct{
		Machine 	model.Machine
		Admin 		[]string
	}
}

func RomGet(c *gin.Context) (result RomResponse,err error) {
	code := c.Param("code")

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionRom).Pipe([]bson.M{
		{"$unwind": "$rom"},
		{"$match": bson.M{"rom.code": code}},
		{"$project": bson.M{"rom": 1,"_id": 0}},
	}).One(&result)

	return result, err
}

func MachineGet(c *gin.Context) (result MachineResponse,err error) {
	code := c.Param("code")
	ip := c.Param("ip")

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionRom).Pipe([]bson.M{
		{"$unwind": "$rom"},
		{"$unwind": "$rom.machine"},
		{"$match": bson.M{"rom.machine.ip": ip, "rom.code": code}},
		{"$project": bson.M{"rom": bson.M{"machine": 1, "admin": 1},"_id": 0}},
	}).One(&result)

	return result, err
}

func OpenRom(c *gin.Context) {
	result, err := RomGet(c)

	session := sessions.Default(c)
	v := session.Get("NowUser")
	user := v.(*util.Claims)
	if user.Auth != "admin" && !util.In_array(user.Username, result.Rom.Admin) {
		c.JSON(e.SUCCESS, gin.H{
			"status" : e.UNAUTHORRIZED,
			"statusText" : e.GetMsg(e.UNAUTHORRIZED),
		})
	} else {
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
}

func CloseRom(c *gin.Context) {
	result, err := RomGet(c)

	session := sessions.Default(c)
	v := session.Get("NowUser")
	user := v.(*util.Claims)
	if user.Auth != "admin" && !util.In_array(user.Username, result.Rom.Admin) {
		c.JSON(e.SUCCESS, gin.H{
			"status" : e.UNAUTHORRIZED,
			"statusText" : e.GetMsg(e.UNAUTHORRIZED),
		})
	} else {
		// Log Rom data for test
		fmt.Println(result.Rom.Machine)

		// TODO: each machine to connect server with telnet for close network

		if err != nil {
			c.Error(err)
		}

		c.JSON(e.SUCCESS, gin.H{
			"status":     e.SUCCESS,
			"statusText": e.GetMsg(e.SUCCESS),
		})
	}
}

func OpenMachine(c *gin.Context) {
	result, err := MachineGet(c)

	session := sessions.Default(c)
	v := session.Get("NowUser")
	user := v.(*util.Claims)
	if user.Auth != "admin" && !util.In_array(user.Username, result.Rom.Admin) {
		c.JSON(e.SUCCESS, gin.H{
			"status" : e.UNAUTHORRIZED,
			"statusText" : e.GetMsg(e.UNAUTHORRIZED),
		})
	} else {
		// Log Rom data for test
		fmt.Println(result.Rom.Machine)

		// TODO: each machine to connect server with telnet for close network

		if err != nil {
			c.Error(err)
		}

		c.JSON(e.SUCCESS, gin.H{
			"status":     e.SUCCESS,
			"statusText": e.GetMsg(e.SUCCESS),
		})
	}
}

func CloseMachine(c *gin.Context) {
	result, err := MachineGet(c)

	session := sessions.Default(c)
	v := session.Get("NowUser")
	user := v.(*util.Claims)
	if user.Auth != "admin" && !util.In_array(user.Username, result.Rom.Admin) {
		c.JSON(e.SUCCESS, gin.H{
			"status" : e.UNAUTHORRIZED,
			"statusText" : e.GetMsg(e.UNAUTHORRIZED),
		})
	} else {
		// Log Rom data for test
		fmt.Println(result.Rom.Machine)

		// TODO: each machine to connect server with telnet for close network

		if err != nil {
			c.Error(err)
		}

		c.JSON(e.SUCCESS, gin.H{
			"status":     e.SUCCESS,
			"statusText": e.GetMsg(e.SUCCESS),
		})
	}
}
