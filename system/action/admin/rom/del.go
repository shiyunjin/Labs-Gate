package rom

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Del(c *gin.Context) {
	code := c.Param("code")
	ip := c.Param("ip")

	db := c.MustGet("db").(*mgo.Database)

	err := db.C(model.CollectionRom).Update(bson.M{
		"rom.code": code,
	},bson.M{
		"$pull": bson.M{
			"rom.$.machine": bson.M{
				"ip": ip,
			},
		},
	})

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
	})
}