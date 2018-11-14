package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/e"
	"github.com/shiyunjin/Labs-Gate/system/util"
)

func Profile(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("NowUser")
	var user *util.Claims

	if v == nil {
		c.JSON(e.SUCCESS, gin.H{
			"name": "???",
			"department": "guest",
			"avatar": "/public/avatar.png",
		})
		return
	} else {
		user = v.(*util.Claims)
	}

	c.JSON(e.SUCCESS, gin.H{
		"name": user.Name,
		"department": user.Auth,
		"avatar": "/public/avatar.png",
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