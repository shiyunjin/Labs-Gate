package device

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	devices := []model.Device{}

	err := db.C(model.CollectionDevice).Find(nil).Sort("name").All(&devices)
	if err != nil {
		c.Error(err)
	}

	var list []gin.H

	for _, device := range devices {
		if device.Vlan == nil {
			device.Vlan = []string{}
		}
		if device.Invalid == nil {
			device.Invalid = []string{}
		}
		tempres := gin.H{
			"id": 		device.Id.Hex(),
			"name": 	device.Name,
			"code": 	device.Code,
			"ip":   	device.Ip,
			"username": device.Username,
			"vlan":		device.Vlan,
			"invalid":  device.Invalid,
		}
		list = append(list, tempres)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : list,
	})
}
