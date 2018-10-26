package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"github.com/shiyunjin/SchoolNetwork/system/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PassRequest struct {
	Oldpassword		string
	Newpassword		string
	Repassword		string
}

func Pass(c *gin.Context) {
	var PassData PassRequest
	err := c.BindJSON(&PassData)

	session := sessions.Default(c)
	user := session.Get("NowUser").(*util.Claims)

	db := c.MustGet("db").(*mgo.Database)

	tempuser := model.User{}
	err = db.C(model.CollectionUser).Find(bson.M{
		"_id": user.Id,
	}).One(&tempuser)

	checkhash := util.HmacSha1(PassData.Oldpassword, tempuser.Salt)

	if tempuser.Hash != checkhash {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.ERROR,
			"statusText":       e.GetMsg(e.ERROR),
		})
		return
	}

	salt, err := password.Generate(32,10,0,false,false)
	if err != nil {
		c.Error(err)
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.ERROR,
			"statusText":       e.GetMsg(e.ERROR),
		})
		return
	}

	hash := util.HmacSha1(PassData.Newpassword, salt)

	err = db.C(model.CollectionUser).Update(bson.M{
		"_id": user.Id,
	}, bson.M{
		"$set": bson.M{
			"hash": hash,
			"salt": salt,
			"updatetime": bson.Now(),
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