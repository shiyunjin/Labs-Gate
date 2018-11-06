package lab

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AddRequest struct {
	Floor	string			`json:"floor"`
	Name 	string			`json:"name"`
	Code 	string			`json:"code"`
	Device 	string			`json:"device"`
	Vlan 	string			`json:"vlan"`
}


func Add(c *gin.Context) {
	var AddData AddRequest
	err := c.BindJSON(&AddData)

	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.INVALID_PARAMS,
			"statusText":       e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionRom).Update(bson.M{
		"_id": bson.ObjectIdHex(AddData.Floor),
	},bson.M{
		"$push":bson.M{
			"rom":bson.M{
				"name":   AddData.Name,
				"code":   AddData.Code,
				"vlan":   AddData.Vlan,
				"device": AddData.Device,
			},
		},
	})

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
		"data":				RomItem{
			Floor: 	AddData.Floor,
			Name: 	AddData.Name,
			Code: 	AddData.Code,
			Device: AddData.Device,
			Vlan: 	AddData.Vlan,
			Machine:"0台机器",
			Admin: 	[]bson.ObjectId{},
		},
	})
}
