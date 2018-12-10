package machine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Add(c *gin.Context){
	code := c.PostForm("code")
	lab := c.PostForm("lab")
	ip := c.PostForm("ip")
	mac := c.PostForm("mac")
	des := c.PostForm("des")
	fmt.Println(code)

	db := c.MustGet("db").(*mgo.Database)

	err := db.C(model.CollectionRom).Update(bson.M{
		"rom.code": lab,
	},bson.M{
		"$push": bson.M{
			"rom.$.machine": bson.M{
				"code": code,
				"ip": ip,
				"mac": mac,
				"des": des,
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
