package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/action/admin/floor"
	"github.com/shiyunjin/SchoolNetwork/system/action/admin/lab"
	"github.com/shiyunjin/SchoolNetwork/system/action/network"
	"github.com/shiyunjin/SchoolNetwork/system/action/rom"
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


			userGroup := api.Group("/user")
			{
				userGroup.POST("/pass", user.Pass)
			}

			api.GET("/roms", rom.List)

			romGroup := api.Group("/rom")
			{
				romGroup.POST("/:code/open", network.OpenRom)
				romGroup.POST("/:code/close", network.CloseRom)

				romGroup.GET("/:code/machine", rom.Machine)
				romGroup.POST("/:code/machine/:ip/open", network.OpenMachine)
				romGroup.POST("/:code/machine/:ip/close", network.CloseMachine)
			}

			api.Use(admin.Need())
			{
				adminUserGroup := api.Group("/user")
				{
					adminUserGroup.GET("/list", AdminUser.List)
					adminUserGroup.POST("/edit", AdminUser.Edit)
					adminUserGroup.POST("/reset", AdminUser.Reset)
					adminUserGroup.POST("/del", AdminUser.Del)
					adminUserGroup.POST("/add", AdminUser.Add)
				}

				adminFloorGroup := api.Group("/floor")
				{
					adminFloorGroup.GET("", 			floor.List)
					adminFloorGroup.POST("/edit", 	floor.Edit)
					adminFloorGroup.POST("/add", 	floor.Add)
					adminFloorGroup.POST("/del", 	floor.Del)
				}

				adminlabGroup := api.Group("/lab")
				{
					adminlabGroup.GET("", 		lab.List)
					adminlabGroup.POST("/add", 	lab.Add)
					adminlabGroup.POST("/del",	lab.Del)
					adminlabGroup.POST("/edit",	lab.Edit)
				}
			}
		}
	}

	return r
}
