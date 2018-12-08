package service

import (
	"errors"
	"fmt"
	"github.com/shiyunjin/Labs-Gate/system/action/network"
	"github.com/shiyunjin/Labs-Gate/system/db"
	"github.com/shiyunjin/Labs-Gate/system/model"
	"github.com/shiyunjin/Labs-Gate/system/service/model"
	"github.com/shiyunjin/go-telnet-cisco"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func ErrorNet(net serviceModel.NetMsg, client *telnet.Client,err error) {
	net.Callback <- err
	_, _ = client.Cmd("exit")
}


func Server(Channel serviceModel.Channel) {

	go Network(Channel)
	//go Bandwidth(Channel)
	//go Status(Channel)
}

//func Bandwidth(Channel serviceModel.Channel) {
//	s := db.Session.Clone()
//	defer s.Close()
//
//	db := s.DB(db.Mongo.Database)
//	for msg := range Channel.Bandwidthch {
//		msg.Callback <- nil
//	}
//}

//func Status(Channel serviceModel.Channel) {
//	for msg := range Channel.StatusCh {
//
//	}
//}

func Network(Channel serviceModel.Channel) {
	s := db.Session.Clone()
	defer s.Close()

	db := s.DB(db.Mongo.Database)
	for net := range Channel.NetworkCh {
		switch net.Type {
		case 1:
			rom := net.Data.(network.RomResponse)
			device := model.Device{}
			err := db.C(model.CollectionDevice).Find(bson.M{"code": rom.Rom.Device}).One(&device)
			if err != nil {
				net.Callback <- err
				return
			}
			client := new(telnet.Client)
			err = client.Connect(device.Ip + ":23")
			if err != nil {
				ErrorNet(net, client, err)
				return
			}
			fmt.Println("connect passed")
			err = client.Login(device.Username, device.Password, device.Super)
			if err != nil {
				ErrorNet(net, client, err)
				return
			}
			fmt.Println("login passed")
			conf,err := client.Cmd("conf t")
			if err != nil {
				ErrorNet(net, client, err)
				return
			}
			if !strings.Contains(conf ,"(config)#") {
				ErrorNet(net, client, errors.New("cant login conf t mode"))
				return
			}
			fmt.Println("conft passed")
			text,err := client.Cmd("interface vlan " + rom.Rom.Vlan)
			if err != nil {
				ErrorNet(net, client, err)
				return
			}
			if strings.Contains(text ,"% Invalid") {
				ErrorNet(net, client, errors.New("invalid rom vlan: cant find"))
				return
			}
			fmt.Println(text)
			text,err = client.Cmd("no ip access-group all_off in")
			if err != nil {
				ErrorNet(net, client, err)
				return
			}
			text,err = client.Cmd("no ip access-group ALL_Open in")
			if err != nil {
				ErrorNet(net, client, err)
				return
			}
			switch net.Open {
			case 1:
				text,err = client.Cmd("ip access-group ALL_Open in")
				if err != nil {
					ErrorNet(net, client, err)
					return
				}
				ErrorNet(net, client, nil)
			case 0:
				text,err = client.Cmd("ip access-group all_off in")
				if err != nil {
					ErrorNet(net, client, err)
					return
				}
				ErrorNet(net, client, nil)
			default:
				ErrorNet(net, client, errors.New("use error open data"))
			}
		case 2:
			machine := net.Data.(network.MachineResponse)
			device := model.Device{}
			err := db.C(model.CollectionDevice).Find(bson.M{"code": machine.Rom.Device}).One(&device)
			if err != nil {
				net.Callback <- err
				return
			}
			switch net.Open {
			case 1:
				net.Callback <- nil
			case 0:
				net.Callback <- nil
			default:
				net.Callback <- errors.New("use error open data")
			}
		default:
			net.Callback <- errors.New("use error type")
		}
	}
}
