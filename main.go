package main

import (
	"github.com/shiyunjin/Labs-Gate/system"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"github.com/shiyunjin/Labs-Gate/system/db"
)

func main() {
	// Init connect to mongodb
	db.Connect();

	r := router.Router()

	// Listen and Server on Port
	r.Run(":" + config.Get("port").(string))
}