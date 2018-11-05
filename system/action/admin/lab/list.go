package lab

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type RomItem struct {
	Floor	string			`json:"floor"`
	Name 	string			`json:"name"`
	Code 	string			`json:"code"`
	Device 	string			`json:"device"`
	Vlan 	string			`json:"vlan"`
	Machine string 			`json:"machine"`
	Admin 	[]bson.ObjectId `json:"admin"`
}

func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roms := []model.Roms{}

	err := db.C(model.CollectionRom).Find(nil).Sort("name").All(&roms)
	if err != nil {
		c.Error(err)
	}

	list := []RomItem{}

	for _, lou := range roms {
		for _, rom := range lou.Rom {
			temprom := RomItem{
				Floor: 	lou.Id.Hex(),
				Name: 	rom.Name,
				Code: 	rom.Code,
				Device: rom.Device,
				Vlan: 	rom.Vlan,
				Machine:strconv.Itoa(len(rom.Machine)) + "台机器",
				Admin: 	rom.Admin,
			}
			list = append(list, temprom)
		}
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : list,
	})
}