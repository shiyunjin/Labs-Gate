package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type AddRequest struct {
	Username	string
	Name		string
	Auth		string
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

	err = db.C(model.CollectionUser).Insert(bson.M{
		"username": AddData.Username,
		"name": AddData.Name,
		"permission": AddData.Auth,
		"createtime": bson.Now(),
		"updatetime": bson.Now(),
	})

	if err != nil {
		c.Error(err)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
	})
}