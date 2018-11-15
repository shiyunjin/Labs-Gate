package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"testing"
)

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.Init()

	router := Router()

	if router == nil {
		t.Fatalf("router create has error")
	}
}
