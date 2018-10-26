package jwt

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/e"
	"github.com/shiyunjin/SchoolNetwork/system/model"
	"github.com/shiyunjin/SchoolNetwork/system/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var claims *util.Claims

		code = e.UNAUTHORRIZED
		token, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(e.SUCCESS, gin.H{
				"status" : code,
				"statusText" : e.GetMsg(code),
				"currentAuthority": "guest",
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
			c.JSON(e.SUCCESS, gin.H{
				"status" : code,
				"statusText" : e.GetMsg(code),
				"currentAuthority": "guest",
			})

			c.Abort()
			return
		}

		db := c.MustGet("db").(*mgo.Database)
		user := model.User{}

		err = db.C(model.CollectionUser).Find(bson.M{
			"_id": claims.Id,
		}).One(&user)

		if user.Hash != claims.Hash {
			c.JSON(e.SUCCESS, gin.H{
				"status" : e.ERROR_AUTH_CHECK_TOKEN_FAIL,
				"statusText" : e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL),
				"currentAuthority": "guest",
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