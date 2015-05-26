package server

import (
	"bufio"
	"fmt"
	"net"
	"object"
	"strconv"
	"time"
)

var _listener net.Listener

//var _mapPlayer map[string]*object.Player   --gai

var _mapPlayer map[*net.Conn]*object.Player

//func GetMapPlayer() map[string]*object.Player { --gai
func GetMapPlayer() map[*net.Conn]*object.Player {
	//func GetMapPlayer() map[net.Conn]*object.Player {
	if _mapPlayer == nil {
		//_mapPlayer = make(map[string]*object.Player, 100)    --gai
		_mapPlayer = make(map[*net.Conn]*object.Player, 100)
	}
	return _mapPlayer
}

func GetLocalServer() net.Listener {
	if _listener == nil {
		var err error
		_listener, err = net.Listen("tcp", "127.0.0.1:9527")
		if err != nil {
			return nil
		}
	}
	return _listener
}

func StartServer() {
	listener := GetLocalServer()
	if listener == nil {
		return
	}
	var conn net.Conn

	for {
		var err error
		conn, err = listener.Accept()
		fmt.Println("收到连接")
		if err != nil {
			continue
		}
		// go func(con net.Conn) {   --gai
		go func(con *net.Conn) {
			//b := make([]byte, 1024)
		L:
			for {
				select {
				case <-time.NewTimer(10 * time.Second).C:
					fmt.Println("连接超市")
					break L
				default:
					fmt.Println("开始接收信息")
					reader := bufio.NewReader(*con)
					b, err := reader.ReadBytes('\n')
					//n, err := (con).Read(b)
					if err != nil {
						fmt.Println("消息错误")
						break L
						continue
					}
					b = b[:len(b)-1] //去除最后一个字符
					//playerId, err1 := strconv.Atoi(string(b[:n]))
					//playerId, err1 := strconv.ParseUint(string(b[:n]), 10, 64)
					playerId, err1 := strconv.ParseUint(string(b), 10, 64)

					if err1 != nil {
						//fmt.Println("转换错误", string(b[:n]))
						fmt.Println("转换错误", string(b))
						break L
						continue
					}
					fmt.Println(playerId)

					player := object.GetGlobalPlayers()[uint64(playerId)]
					if player == nil {
						continue
					}
					//GetMapPlayer()[con.RemoteAddr().String()] = player   --gai
					GetMapPlayer()[con] = player

					//fmt.Println("globalMap Len", len(GetMapPlayer()))

					//player.StartConn()
					fmt.Println("接收完成")
					StartConn(con)
					break L
				}
			}
		}(&conn)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func StartConn(conn *net.Conn) {
	go func(con *net.Conn) {
		fmt.Println("conn准备收到消息：")
		(*con).Write([]byte("\n"))
		reader := bufio.NewReader(*conn)
		for {
			//b := make([]byte, 1024)
			//(con).Read(b)
			b, err := reader.ReadBytes('\n')
			if err != nil {
				continue
			}
			fmt.Println("conn收到消息：", string(b))
			//player := GetMapPlayer()[conn.RemoteAddr().String()]  --gai
			player := GetMapPlayer()[con]
			if player == nil {
				fmt.Println("player错误")
				break
			}
			fmt.Println("发送player消息", string(b))
			player.Mq <- string(b)
		}
	}(conn)
}
