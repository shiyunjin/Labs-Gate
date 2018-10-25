package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/util"
)

func Profile(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("NowUser")
	var user *util.Claims

	if v == nil {
		c.JSON(e.SUCCESS, gin.H{
			"name": "???",
			"department": "guest",
			"avatar": "https://img.alicdn.com/tfs/TB1L6tBXQyWBuNjy0FpXXassXXa-80-80.png",
		})
		return
	} else {
		user = v.(*util.Claims)
	}

	c.JSON(e.SUCCESS, gin.H{
		"name": user.Name,
		"department": user.Auth,
		"avatar": "https://img.alicdn.com/tfs/TB1L6tBXQyWBuNjy0FpXXassXXa-80-80.png",
	})
}

func Authority(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("NowUser")

	var auth string
	if v == nil {
		auth = "guest"
	} else {
		auth = v.(*util.Claims).Auth
	}

	c.JSON(e.SUCCESS, gin.H{
		"status":           e.SUCCESS,
		"statusText":       e.GetMsg(e.SUCCESS),
		"currentAuthority": auth,
	})
}