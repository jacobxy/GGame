package server

import (
	"bufio"
	//"error"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"object"
	"strconv"
	"time"
)

const (
	CHAN_IN_MAX    = 100
	TCP_TIME_LIMIT = 120
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

//第一次接受信息验证玩家身份
func handleClient(con *net.Conn) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("caught panic in handleClient")
		}
	}()

	fmt.Println("开始接收信息")
	reader := bufio.NewReader(*con)
	b, err := reader.ReadBytes('\n')
	//n, err := (con).Read(b)
	if err != nil {
		fmt.Println("消息错误")
		return
	}
	b = b[:len(b)-1] //去除最后一个字符
	playerId, err1 := strconv.ParseUint(string(b), 10, 64)

	if err1 != nil {
		//fmt.Println("转换错误", string(b[:n]))
		fmt.Println("转换错误", string(b))
		return
	}
	fmt.Println(playerId)

	player := object.GetGlobalPlayers()[uint64(playerId)]
	if player == nil {
		return
	}
	//GetMapPlayer()[con.RemoteAddr().String()] = player   --gai
	GetMapPlayer()[con] = player

	//player.StartConn()
	fmt.Println("接收完成")
	StartConn(con)
	//StartServer2(con)
}

// 模仿版
func StartServer2() {
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
		go handleClient2(&conn)
	}
}
func handleClient2(conn *net.Conn) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("error happen in handleClient2")
		}
	}()

	in := make(chan []byte, CHAN_IN_MAX)

	out := NewConnWrite(*conn)

	fmt.Println("handclient2")
	LineInAndOut(in, out)

	header := make([]byte, 2)

	go func(con net.Conn) {
		//L:
		for {
			//con.SetReadDeadline(time.Now().Add(10 * time.Second))

			n, err := io.ReadFull(con, header)
			fmt.Println("开始接收信息")
			if err != nil {
				fmt.Println("error receiving header, bytes:", n, "reason:", err)
				break
			}

			size := binary.BigEndian.Uint16(header)
			fmt.Println(size)

			data := make([]byte, size)

			_, err = io.ReadFull(con, data)

			select {
			case in <- data:
			case <-time.After(200 * time.Second):
			}
		}
	}(*conn)
}

func LineInAndOut(in chan []byte, out *ConnWrite) {
	go func() {
		for {
			select {
			case _, ok := <-in:
				if !ok {
				}
			}
		}
	}()
}
