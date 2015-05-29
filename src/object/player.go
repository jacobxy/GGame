package object

import (
	//"time"
	//"encoding/json"
	"fmt"
	//"net"
	"strings"
)

type Player struct {
	Id         uint64
	Name       string
	Level      uint8
	MyV        *Var
	Mq         chan string
	MyFighters map[uint32]*Fighter
	clan       *Clan
}

type JsonStruct struct {
	method string `json:"Method"`
	param  string `json:"Param"`
}

var _function map[string]func([]string) uint8

func NewPlayer() *Player {
	m := make(chan string, 0)
	myVar := &Var{MyVar: map[uint32]uint32{}, MyVarOver: map[uint32]uint32{}}
	player := &Player{Id: 0, Name: "", Level: 0, MyV: myVar, Mq: m, MyFighters: make(map[uint32]*Fighter)}
	go func(pl *Player) {
		for {
			msg := <-pl.Mq
			fmt.Println("player收到消息：", msg)
			tokens := strings.Split(msg, ":")
			if len(tokens) == 1 {
				fmt.Println(tokens[0])
			} else if len(tokens) == 2 {
				pl.Handler(tokens[0], tokens[1])
			} else {
				continue
			}
		}
	}(player)

	return player
}

func (pl *Player) SetClan(cl *Clan) {
	pl.clan = cl
}

func (pl *Player) LoadPlayerInfo(id uint64, name string, level uint8) {
	pl.Id = id
	pl.Name = name
	pl.Level = level
}

func (pl *Player) SendMessage(param string) bool {
	fmt.Println("SendMessage :", param)
	return true
}

func (pl *Player) GetMapFunction() map[string]func(string) bool {
	return map[string]func(string) bool{
		"send":     pl.SendMessage,
		"get":      pl.SendMessage,
		"AddMoney": pl.SendMessage,
		"AddVar":   pl.SendMessage,
	}
}

func (pl *Player) Handler(method string, param string) {
	if handler, ok := pl.GetMapFunction()[method]; ok {
		ret := handler(param)
		if ret {
			fmt.Println(method, param)
		}
	} else {
		fmt.Println("Unknow method")
	}
}

func (pl *Player) GetFighters() map[uint32]*Fighter {
	return pl.MyFighters
}