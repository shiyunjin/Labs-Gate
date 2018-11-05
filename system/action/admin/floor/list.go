package floor

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roms := []model.Roms{}

	err := db.C(model.CollectionRom).Find(nil).Select(bson.M{"_id": 1, "name": 1}).Sort("name").All(&roms)
	if err != nil {
		c.Error(err)
	}

	list := make(map[string]string)

	for _, lou := range roms {
		tempkey := lou.Id.Hex()
		list[tempkey] = lou.Name
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : list,
	})
}