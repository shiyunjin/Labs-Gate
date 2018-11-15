package main

import (
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func TestLocal(t *testing.T) {
	gin.SetMode(gin.TestMode)

	go func() {
		main()
	}()

	time.Sleep(3 * time.Second)
}