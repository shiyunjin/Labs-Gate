package rom

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type TabResponse struct {
	Tab			string	`json:"tab"`
	Key 		int		`json:"key"`
}

type DataResponse struct {
	Tabs 		[]TabResponse	`json:"tabs"`
	DataSource	[][]RomResponse`json:"dataSource"`
}

type RomResponse struct {
	Name		string	`json:"name"`
	Code 		string	`json:"code"`
	Desc 		string	`json:"desc"`
	Acl 		bool	`json:"acl"`
}

func List(c *gin.Context) {
	//session := sessions.Default(c)
	//user := session.Get("NowUser").(*util.Claims)

	db := c.MustGet("db").(*mgo.Database)
	roms := []model.Roms{}
	// TODO: 目前是显示所有机器，需要按权限显示
	err := db.C(model.CollectionRom).Find(nil).Sort("name").All(&roms)
	if err != nil {
		c.Error(err)
	}

	var data DataResponse

	for index, lou := range roms {
		tempTab := TabResponse{
			Tab: lou.Name,
			Key: index,
		}
		data.Tabs = append(data.Tabs, tempTab)
		tempList := []RomResponse{}
		for _, rom := range lou.Rom {
			tempRom := RomResponse{
				Name: rom.Name,
				Code: rom.Code,
				Desc: "本实验室机器数量：" + strconv.Itoa(len(rom.Machine)) + "台",
				Acl:  true, // TODO: add status
			}
			tempList = append(tempList, tempRom)
		}
		data.DataSource = append(data.DataSource, tempList)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : data,
	})
}

type MachineResponse struct {
	Rom		model.Rom
}

type MachineData struct {
	Ip 		string 	`json:"ip"`
	Mac 	string	`json:"mac"`
	Des 	string	`json:"des"`
	Status 	string	`json:"status"`
}

func Machine(c *gin.Context) {
	code := c.Param("code")

	db := c.MustGet("db").(*mgo.Database)
	rom := MachineResponse{}
	// TODO: 目前是显示所有机器，需要按权限显示

	err := db.C(model.CollectionRom).Pipe([]bson.M{
		{"$unwind": "$rom"},
		{"$match": bson.M{"rom.code": code}},
		{"$project": bson.M{"rom": 1,"_id": 0}},
	}).One(&rom)

	if err != nil {
		c.Error(err)
	}

	var data []MachineData

	for _, machine := range rom.Rom.Machine {
		tempMachine := MachineData{
			Ip:		machine.Ip,
			Mac: 	machine.Mac,
			Des: 	machine.Des,
			Status: "OPEN",
		}
		data = append(data, tempMachine)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : data,
	})
}