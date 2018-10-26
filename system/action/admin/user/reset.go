package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"github.com/shiyunjin/SchoolNetwork/system/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ResetRequest struct {
	Id		string
}

func Reset(c *gin.Context) {
	var ResetData ResetRequest
	err := c.BindJSON(&ResetData)

	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.INVALID_PARAMS,
			"statusText":       e.GetMsg(e.INVALID_PARAMS),
		})
		return
	}

	salt, err := password.Generate(32,10,0,false,false)
	if err != nil {
		c.Error(err)
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.ERROR,
			"statusText":       e.GetMsg(e.ERROR),
			"data":				e.GetMsg(e.ERROR),
		})
		return
	}
	pass, err := password.Generate(8,3,0,false,false)
	if err != nil {
		c.Error(err)
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.ERROR,
			"statusText":       e.GetMsg(e.ERROR),
			"data":				e.GetMsg(e.ERROR),
		})
		return
	}

	hash := util.HmacSha1(pass, salt)

	db := c.MustGet("db").(*mgo.Database)

	err = db.C(model.CollectionUser).Update(bson.M{
		"_id": bson.ObjectIdHex(ResetData.Id),
	}, bson.M{
		"$set": bson.M{
			"hash": hash,
			"salt": salt,
		},
	})

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
		"data":				pass,
	})
}