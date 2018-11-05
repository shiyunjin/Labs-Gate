package floor

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EditRequest struct {
	Id		string
	Name	string
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
		"_id": bson.ObjectIdHex(EditData.Id),
	}, bson.M{
		"$set": bson.M{
			"name": EditData.Name,
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