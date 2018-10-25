package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/util"
	"net/http"
)

func Need() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("NowUser").(*util.Claims)

		if user.Auth != "admin" {
			code := e.UNAUTHORRIZED

			c.JSON(http.StatusUnauthorized, gin.H{
				"status" : code,
				"statusText" : e.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}