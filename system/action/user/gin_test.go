package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"github.com/shiyunjin/Labs-Gate/system/db"
	"github.com/shiyunjin/Labs-Gate/system/middlewares"
	"github.com/shiyunjin/Labs-Gate/system/util"
)

func fuckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := &util.Claims{
			Id:       "5bd00e64b847d59e683f2024",
			Name:     "syj",
			Username: "syj",
			Auth:     "admin",
			Hash:     "testhash",
		}

		session := sessions.Default(c)
		session.Set("NowUser", claims)

		c.Next()
	}
}

func testGinWithLogin() (server *gin.Engine) {
	gin.SetMode(gin.TestMode)

	config.Init()
	util.JwtInit()

	db.Connect()

	server = gin.New()
	server.Use(middlewares.Connect)

	// Support session
	store := cookie.NewStore([]byte(config.Get("secret").(string)))
	server.Use(sessions.Sessions("SESSION", store))
	server.Use(fuckJWT())

	return server
}


func testGin() (server *gin.Engine) {
	gin.SetMode(gin.TestMode)

	config.Init()
	util.JwtInit()

	db.Connect()

	server = gin.New()
	server.Use(middlewares.Connect)

	// Support session
	store := cookie.NewStore([]byte(config.Get("secret").(string)))
	server.Use(sessions.Sessions("SESSION", store))

	return server
}
