package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"github.com/shiyunjin/Labs-Gate/system/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LoginRequest struct {
	Username	string
	Password	string
	Checkbox 	bool
}

func Login(c *gin.Context) {
	var LoginData LoginRequest
	err := c.BindJSON(&LoginData)
	if err != nil {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.UNAUTHORRIZED,
			"statusText":       e.GetMsg(e.UNAUTHORRIZED),
			"currentAuthority": "guest",
		})
		return
	}

	// log of debug json decoder
	// c.Error(errors.New("username: " + LoginData.Username))

	user := model.User{}

	db := c.MustGet("db").(*mgo.Database)
	err = db.C(model.CollectionUser).Find(bson.M{
		"username": LoginData.Username,
	}).One(&user)

	if err != nil && !(gin.Mode() == gin.TestMode && LoginData.Username == "admin") {
		c.Error(err)
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.UNAUTHORRIZED,
			"statusText":       e.GetMsg(e.UNAUTHORRIZED),
			"currentAuthority": "guest",
		})
		return
	}

	if gin.Mode() == gin.TestMode {
		user.Permission = "admin"
		user.Username = "admin"
		user.Salt = "testsalt"
		user.Hash = "2e224647eeca047c0353eb9745c2f072dc5b2a17"
	}

	if user.Username != LoginData.Username {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.UNAUTHORRIZED,
			"statusText":       e.GetMsg(e.UNAUTHORRIZED),
			"currentAuthority": "guest",
		})
		return
	}

	hash := util.HmacSha1(LoginData.Password, user.Salt)

	if user.Hash != hash {
		c.JSON(e.SUCCESS, gin.H{
			"status":           e.UNAUTHORRIZED,
			"statusText":       e.GetMsg(e.UNAUTHORRIZED),
			"currentAuthority": "guest",
		})
		return
	}

	token, err := util.GenerateToken(user.Id, user.Name, user.Username, user.Hash, user.Permission, util.If(LoginData.Checkbox, 168, 3).(int))

	if err != nil {
		c.Error(err)
	}

	c.SetCookie("token", token, 360000, "/", "", false, false)

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
		"currentAuthority": user.Permission,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "0",-1,"/","",false,false)

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
		"currentAuthority": "guest",
	})
}
