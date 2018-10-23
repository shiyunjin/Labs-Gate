package main

import (
	"github.com/shiyunjin/SchoolNetwork/system"
	"github.com/shiyunjin/SchoolNetwork/system/config"
	"github.com/shiyunjin/SchoolNetwork/system/db"
)

func main() {
	r := router.Router()

	// Init connect to mongodb
	db.Connect();

	// Listen and Server on Port
	r.Run(":" + config.Get("port").(string))
}