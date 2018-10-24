package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/SchoolNetwork/system/action/user"
	"github.com/shiyunjin/SchoolNetwork/system/middlewares"
)

func Router() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(middlewares.Connect)

	// Serve the frontend
	r.Use(static.Serve("/", static.LocalFile("system/view/build", true)))

	//API Serve
	api := r.Group("/api/v1")
	{
		api.POST("/login", user.Login)
		api.POST("/logout", user.Logout)

		api.GET("/profile", user.Profile)
	}

	return r
}
