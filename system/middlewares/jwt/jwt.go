package jwt

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *util.Claims

		code = e.UNAUTHORRIZED
		token, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status" : code,
				"statusText" : e.GetMsg(code),
			})

			c.Abort()
			return
		}
		code = e.SUCCESS
		if token.Value == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err = util.ParseToken(token.Value)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		session := sessions.Default(c)
		session.Set("NowUser", claims)
		session.Save()

		c.Next()
	}
}