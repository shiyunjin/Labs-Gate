package main

import (
	"github.com/shiyunjin/SchoolNetwork/system"
	"github.com/shiyunjin/SchoolNetwork/system/config"
	"github.com/shiyunjin/SchoolNetwork/system/db"
)

func main() {
	// Init connect to mongodb
	db.Connect();

	r := router.Router()

	// Listen and Server on Port
	r.Run(":" + config.Get("port").(string))
}