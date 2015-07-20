package main

import (
	//"bufio"
	"fmt"
	"net"
	. "stream"
	//"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9527")
	CheckErr(err)
	st := GetStream()
	st.WriteRawBytes([]byte("844424930131971"))
	//st.WriteString("844424930131971")
	fmt.Println(st.Data())

	st1 := GetStream()
	st1.WriteBytes(st.Data())

	_, err1 := conn.Write(st1.Data())
	CheckErr(err1)

	//reader := bufio.NewReader(conn)

	//reader.ReadBytes('\n')

	//	_, err2 := conn.Write([]byte("send:my name is libo\n"))
	//	CheckErr(err2)
	//	_, err3 := conn.Write([]byte("get:my name is libo\n"))
	//	CheckErr(err3)
	//	time.Sleep(20 * time.Second)
	//	_, err4 := conn.Write([]byte("send:my name is libo\n"))
	//	CheckErr(err4)
	//	conn.Close()
	//	fmt.Println("Over")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
