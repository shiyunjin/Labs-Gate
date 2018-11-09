package device

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type InterfaceRequest struct {
	Id		 string
	Vlan	 []string
	Invalid  []string
}


func Interface(c *gin.Context) {
	var InterfaceData InterfaceRequest
	err := c.BindJSON(&InterfaceData)

	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.INVALID_PARAMS,
			"statusText":       e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionDevice).Update(bson.M{
		"_id": bson.ObjectIdHex(InterfaceData.Id),
	}, bson.M{
		"$set": bson.M{
			"vlan": 	InterfaceData.Vlan,
			"invalid":	InterfaceData.Invalid,
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
