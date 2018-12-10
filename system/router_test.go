package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"github.com/shiyunjin/Labs-Gate/system/service/model"
	"testing"
)

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.Init()

	Channel := serviceModel.Channel{}

	router := Router(Channel)

	if router == nil {
		t.Fatalf("router create has error")
	}
}
