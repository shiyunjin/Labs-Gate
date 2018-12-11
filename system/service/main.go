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
	"regexp"
	"strconv"
	"strings"
)


func Server(Channel serviceModel.Channel) {
	go Network(Channel)
	go Bandwidth(Channel)
	//go Status(Channel)
}

func Bandwidth(Channel serviceModel.Channel) {
	s := db.Session.Clone()
	defer s.Close()

	db := s.DB(db.Mongo.Database)

	for msg := range Channel.Bandwidthch {
		switch msg.Type {
		case 1:
			fmt.Println(msg.Device)

			device := model.Device{}
			err := db.C(model.CollectionDevice).Find(bson.M{"code": msg.Device}).One(&device)
			if err != nil {
				msg.Callback <- serviceModel.BandwidthCall{Err: err, Data: nil}
				continue
			}
			client := new(telnet.Client)
			err = client.Connect(device.Ip + ":23")
			if err != nil {
				ErrorBandwidth(msg, client, nil, err)
				continue
			}
			fmt.Println("connect passed")
			err = client.Login(device.Username, device.Password, device.Super)
			if err != nil {
				ErrorBandwidth(msg, client, nil, err)
				continue
			}
			fmt.Println("login passed")
			text,err := client.Cmd("show interfaces counters summary")
			if err != nil {
				ErrorBandwidth(msg, client, nil, err)
				continue
			}
			nowRegexp := regexp.MustCompile("(G|F)[^\r\n]{1,}")
			interNameRegexp := regexp.MustCompile("((G|F)[\\S]{1,})")
			interBandRegexp := regexp.MustCompile("([\\s]([\\d]{1,}))")
			inter := nowRegexp.FindAllStringSubmatch(text,-1)
			result := make(map[string][]int64)
			for _,one := range inter {
				interName := interNameRegexp.FindStringSubmatch(one[0])
				interBand := interBandRegexp.FindStringSubmatch(one[0])
				temp, _ := strconv.ParseInt(strings.TrimSpace(interBand[0]),10,64)
				result[interName[0]] = append(result[interName[0]],temp)
			}
			ErrorBandwidth(msg, client, result, nil)
		default:
			msg.Callback <- serviceModel.BandwidthCall{Data: nil, Err: errors.New("use error type")}
		}
	}
}

func ErrorBandwidth(msg serviceModel.BandwidthMsg, client *telnet.Client, data interface{}, err error) {
	msg.Callback <- serviceModel.BandwidthCall{Data: data, Err: err}
	_, _ = client.Cmd("exit")
}



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
				continue
			}
			client := new(telnet.Client)
			err = client.Connect(device.Ip + ":23")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			fmt.Println("connect passed")
			err = client.Login(device.Username, device.Password, device.Super)
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			fmt.Println("login passed")
			conf,err := client.Cmd("conf t")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			if !strings.Contains(conf ,"(config)#") {
				ErrorNet(net, client, errors.New("cant login conf t mode"))
				continue
			}
			fmt.Println("conft passed")
			text,err := client.Cmd("interface vlan " + rom.Rom.Vlan)
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			if strings.Contains(text ,"% Invalid") {
				ErrorNet(net, client, errors.New("invalid rom vlan: cant find"))
				continue
			}
			fmt.Println(text)
			text,err = client.Cmd("no ip access-group acl_" + rom.Rom.Vlan + " in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("no ip access-group all_off_hosts in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("no ip access-group all_off in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("no ip access-group ALL_Open in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("exit")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			switch net.Open {
			case 1:
				text,err = client.Cmd("no ip access-list extended acl_" + rom.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("ip access-list extended acl_" + rom.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit ip any any")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("exit")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("interface vlan " + rom.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("ip access-group acl_" + rom.Rom.Vlan + " in")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("exit")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				ErrorNet(net, client, nil)
			case 0:
				text,err = client.Cmd("no ip access-list extended acl_" + rom.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("ip access-list extended acl_" + rom.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit ip 192.168.0.220 0.0.255.3 any")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit ip any 172.16.0.0 0.0.255.255")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit udp any any range bootps bootpc")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit ip any 192.168.100.0 0.0.0.255")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit ip 192.168.100.0 0.0.0.255 any")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit tcp any host 1.1.1.8 eq telnet")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit udp any any eq domain")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("permit tcp any any eq domain")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("exit")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("interface vlan " + rom.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("ip access-group acl_" + rom.Rom.Vlan + " in")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("exit")
				if err != nil {
					ErrorNet(net, client, err)
					continue
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
				continue
			}
			client := new(telnet.Client)
			err = client.Connect(device.Ip + ":23")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			fmt.Println("connect passed")
			err = client.Login(device.Username, device.Password, device.Super)
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			fmt.Println("login passed")
			conf,err := client.Cmd("conf t")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			if !strings.Contains(conf ,"(config)#") {
				ErrorNet(net, client, errors.New("cant login conf t mode"))
				continue
			}
			fmt.Println("conft passed")
			text,err := client.Cmd("interface vlan " + machine.Rom.Vlan)
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			if strings.Contains(text ,"% Invalid") {
				ErrorNet(net, client, errors.New("invalid rom vlan: cant find"))
				continue
			}
			fmt.Println(text)
			text,err = client.Cmd("no ip access-group acl_" + machine.Rom.Vlan + " in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("no ip access-group all_off_hosts in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("no ip access-group all_off in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("no ip access-group ALL_Open in")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			text,err = client.Cmd("exit")
			if err != nil {
				ErrorNet(net, client, err)
				continue
			}
			switch net.Open {
			case 1:
				text,err = client.Cmd("ip access-list extended acl_" + machine.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				switch machine.Rom.Acl {
				case true:
					text,err = client.Cmd("permit ip host " + machine.Rom.Machine.Ip + " any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("exit")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
				default:
					text,err = client.Cmd("no permit ip any any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("no deny ip host " + machine.Rom.Machine.Ip + " any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("permit ip any any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("exit")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
				}
				text,err = client.Cmd("interface vlan " + machine.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("ip access-group acl_" + machine.Rom.Vlan + " in")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("exit")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				ErrorNet(net, client, nil)
			case 0:
				text,err = client.Cmd("ip access-list extended acl_" + machine.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				switch machine.Rom.Acl {
				case true:
					text,err = client.Cmd("no permit ip host " + machine.Rom.Machine.Ip + " any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("exit")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
				default:
					text,err = client.Cmd("no permit ip any any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("deny ip host " + machine.Rom.Machine.Ip + " any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("permit ip any any")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
					text,err = client.Cmd("exit")
					if err != nil {
						ErrorNet(net, client, err)
						continue
					}
				}
				text,err = client.Cmd("interface vlan " + machine.Rom.Vlan)
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("ip access-group acl_" + machine.Rom.Vlan + " in")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				text,err = client.Cmd("exit")
				if err != nil {
					ErrorNet(net, client, err)
					continue
				}
				ErrorNet(net, client, nil)
			default:
				ErrorNet(net, client, errors.New("use error open data"))
			}
		default:
			net.Callback <- errors.New("use error type")
		}
	}
}

func ErrorNet(net serviceModel.NetMsg, client *telnet.Client,err error) {
	net.Callback <- err
	_, _ = client.Cmd("exit")
}
