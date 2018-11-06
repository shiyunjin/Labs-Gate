package lab

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DelRequest struct {
	Floor		string
	Code 		string
}

func Del(c *gin.Context) {
	var DelData DelRequest
	err := c.BindJSON(&DelData)

	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.INVALID_PARAMS,
			"statusText":       e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionRom).Update(bson.M{
		"_id": bson.ObjectIdHex(DelData.Floor),
	},bson.M{
		"$pull": bson.M{
			"rom": bson.M{
				"code": DelData.Code,
			},
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
