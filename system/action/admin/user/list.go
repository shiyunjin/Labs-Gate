package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"gopkg.in/mgo.v2"
)


type ListResponse struct {
	Id			string	`json:"id"`
	Username	string	`json:"username"`
	Name		string	`json:"name"`
	Auth 		string	`json:"auth"`
	Date 		string	`json:"date"`
}


func List(c *gin.Context) {
	//session := sessions.Default(c)
	//user := session.Get("NowUser").(*util.Claims)

	db := c.MustGet("db").(*mgo.Database)
	users := []model.User{}
	err := db.C(model.CollectionUser).Find(nil).Sort("createtime").All(&users)
	if err != nil {
		c.Error(err)
	}

	all := []ListResponse{}
	user := []ListResponse{}
	admin := []ListResponse{}

	for _, row := range users {
		tempUser := ListResponse{
			Id:			row.Id.Hex(),
			Username:	row.Username,
			Name:		row.Name,
			Auth: 		row.Permission,
			Date: 		row.Createtime.Format("2006-01-02 15:04:05"),
		}
		all = append(all, tempUser)
		if row.Permission == "admin" {
			user = append(user, tempUser)
		} else {
			admin = append(admin, tempUser)
		}
	}
	c.JSON(e.SUCCESS, gin.H{
		"status" : e.SUCCESS,
		"statusText" : e.GetMsg(e.SUCCESS),
		"data" : gin.H{
			"all": all,
			"user": user,
			"admin": admin,
		},
	})

}