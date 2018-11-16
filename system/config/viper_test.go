package config

import (
	"fmt"
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

func TestViperPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("viper init panic success", r)
		}
	}()
	gin.SetMode(gin.DebugMode)

	Init()

	Set("testviper","145214452")

	if Get("testviper").(string) != "145214452" {
		t.Fatalf("viper has error")
	}

	t.Errorf("viper did not panic")
}