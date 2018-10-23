package router

import (
	"gitee.com/shiyunjin/SchoolNetwork/system/action"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Serve the frontend
	r.Use(static.Serve("/", static.LocalFile("system/view/SchoolNetworkUI/build", true)))

	//API Serve
	api := r.Group("/api")
	{
		api.GET("/profile", action.Exp)
	}

	return r
}
