package lab

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AdminRequest struct {
	Floor	string			`json:"floor"`
	Code 	string			`json:"code"`
	Admin 	[]string		`json:"admin"`
}

func Admin(c *gin.Context) {
	var AdminData AdminRequest
	err := c.BindJSON(&AdminData)

	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.INVALID_PARAMS,
			"statusText":       e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionRom).Update(bson.M{
		"rom.code":  AdminData.Code,
		"_id": bson.ObjectIdHex(AdminData.Floor),
	}, bson.M{
		"$set": bson.M{
			"rom.$.admin": AdminData.Admin,
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
