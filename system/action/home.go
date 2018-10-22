package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.tmpl", gin.H{
		"title": "Main website",
	})
}