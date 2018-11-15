package db

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"testing"
)

func TestConnect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.Init()

	Connect()
}
