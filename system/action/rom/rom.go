package rom

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"github.com/shiyunjin/Labs-Gate/system/util"
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

	// 管理员显示所有机器，其他用户需要按权限显示
	session := sessions.Default(c)
	v := session.Get("NowUser")
	user := v.(*util.Claims)
	var match bson.M
	if user.Auth == "admin" {
		match = nil
	} else {
		match = bson.M{"rom.admin": user.Username}
	}

	err := db.C(model.CollectionRom).Find(match).Sort("name").All(&roms)
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
			if user.Auth == "admin" || util.In_array(user.Username, rom.Admin) {
				tempRom := RomResponse{
					Name: rom.Name,
					Code: rom.Code,
					Desc: "本实验室机器数量：" + strconv.Itoa(len(rom.Machine)) + "台",
					Acl:  rom.Acl,
				}
				tempList = append(tempList, tempRom)
			}
		}
		data.DataSource = append(data.DataSource, tempList)
	}

	if len(data.Tabs) == 0 {
		data.Tabs = []TabResponse{}
		data.DataSource = [][]RomResponse{}
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

	// 管理员显示所有机器，其他用户需要按权限显示
	session := sessions.Default(c)
	v := session.Get("NowUser")
	user := v.(*util.Claims)
	var match bson.M
	if user.Auth == "admin" {
		match = bson.M{"rom.code": code}
	} else {
		match = bson.M{"rom.code": code, "rom.admin": user.Username}
	}

	err := db.C(model.CollectionRom).Pipe([]bson.M{
		{"$unwind": "$rom"},
		{"$match": match},
		{"$project": bson.M{"rom": 1,"_id": 0}},
	}).One(&rom)

	if err != nil {
		c.Error(err)
	}

	data := []MachineData{}

	for _, machine := range rom.Rom.Machine {
		tempMachine := MachineData{
			Ip:		machine.Ip,
			Mac: 	machine.Mac,
			Des: 	machine.Des,
			Status: util.If(machine.Acl, "CLOSE", "OPEN").(string),
		}
		data = append(data, tempMachine)
	}

	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"name" : rom.Rom.Name,
		"data" : data,
	})
}