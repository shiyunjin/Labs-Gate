package user

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":           "401",
		"statusText":       "unauthorized",
		"currentAuthority": "guest",
	})
}
