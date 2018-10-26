package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/action/user"
	AdminUser "github.com/shiyunjin/SchoolNetwork/system/action/admin/user"
	"github.com/shiyunjin/SchoolNetwork/system/config"
	"github.com/shiyunjin/SchoolNetwork/system/middlewares"
	"github.com/shiyunjin/SchoolNetwork/system/middlewares/admin"
	"github.com/shiyunjin/SchoolNetwork/system/middlewares/jwt"
)

func Router() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(middlewares.Connect)

	// Support session
	store := cookie.NewStore([]byte(config.Get("secret").(string)))
	r.Use(sessions.Sessions("SESSION", store))

	// Serve the frontend
	r.Use(static.Serve("/", static.LocalFile("system/view/build", true)))

	//API Serve
	api := r.Group("/api/v1")
	{
		api.POST("/login", user.Login)
		api.POST("/logout", user.Logout)

		api.Use(jwt.JWT())
		{
			api.GET("/profile", user.Profile)
			api.GET("/authority", user.Authority)

			api.Use(admin.Need())
			{
				userGroup := api.Group("/user")
				{
					userGroup.GET("/list", AdminUser.List)
					userGroup.POST("/edit", AdminUser.Edit)
					userGroup.POST("/reset", AdminUser.Reset)
					userGroup.POST("/del", AdminUser.Del)
				}
			}
		}
	}

	return r
}
