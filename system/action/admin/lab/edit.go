package lab

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EditRequest struct {
	Floor	string			`json:"floor"`
	Name 	string			`json:"name"`
	Code 	string			`json:"code"`
	Device 	string			`json:"device"`
	Vlan 	string			`json:"vlan"`
}

func Edit(c *gin.Context) {
	var EditData EditRequest
	err := c.BindJSON(&EditData)

	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.INVALID_PARAMS,
			"statusText":       e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionRom).Update(bson.M{
		"rom.code":  EditData.Code,
		"_id": bson.ObjectIdHex(EditData.Floor),
	}, bson.M{
		"$set": bson.M{
			"rom.$.name": EditData.Name,
			"rom.$.vlan": EditData.Vlan,
			"rom.$.device": EditData.Device,
		},
	})

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
	})
}
