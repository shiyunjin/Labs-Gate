package main

import (
	"github.com/shiyunjin/Labs-Gate/system"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"github.com/shiyunjin/Labs-Gate/system/db"
	"github.com/shiyunjin/Labs-Gate/system/util"
)

func main() {
	// Loding config
	config.Init()

	// Init connect to mongodb
	db.Connect()

	// Init Jwt
	util.JwtInit()

	r := router.Router()

	// Listen and Server on Port
	r.Run(":" + config.Get("port").(string))
}