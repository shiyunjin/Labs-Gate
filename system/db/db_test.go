package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"testing"
)

func TestConnect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.Init()

	Connect()
}

func TestConnectPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("mongodb connect panic success", r)
		}
	}()
	gin.SetMode(gin.TestMode)

	config.Init()

	config.Set("mongodb.username", "failuser")
	config.Set("mongodb.password", "failuser")

	Connect()
	t.Errorf("mongodb did not panic")
}
