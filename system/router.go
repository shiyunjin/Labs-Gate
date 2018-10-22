package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Serve the frontend
	r.Use(static.Serve("/", static.LocalFile("system/view/build", true)))


	return r
}
