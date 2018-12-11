package main

import (
	"github.com/shiyunjin/Labs-Gate/system"
	"github.com/shiyunjin/Labs-Gate/system/config"
	"github.com/shiyunjin/Labs-Gate/system/db"
	"github.com/shiyunjin/Labs-Gate/system/service"
	"github.com/shiyunjin/Labs-Gate/system/service/model"
	"github.com/shiyunjin/Labs-Gate/system/util"
)

func main() {

	// make chan
	Channel := serviceModel.Channel{}

	Channel.NetworkCh = make(chan serviceModel.NetMsg)
	defer close(Channel.NetworkCh)

	Channel.Bandwidthch = make(chan serviceModel.BandwidthMsg)
	defer close(Channel.Bandwidthch)

	// Loding config
	config.Init()

	// Init connect to mongodb
	db.Connect()

	// Init Jwt
	util.JwtInit()

	// Create Service
	go service.Server(Channel)

	r := router.Router(Channel)

	// Listen and Server on Port
	r.Run(":" + config.Get("port").(string))
}

