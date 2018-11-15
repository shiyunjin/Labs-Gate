package config

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestViper(t *testing.T) {
	gin.SetMode(gin.TestMode)

	Init()

	Set("testviper","145214452")

	if Get("testviper").(string) != "145214452" {
		t.Fatalf("viper has error")
	}
}
