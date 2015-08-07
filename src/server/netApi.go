package server

import (
	//"net"
	//. "object"
	. "stream"
)

var NetHandler map[uint16]func(*ConnWrite, *Stream) []byte

func init() {
	NetHandler = map[uint16]func(*ConnWrite, *Stream) []byte{
		1: handleLogin,
	}
}

func handleLogin(wt *ConnWrite, st *Stream) []byte {
	account, err1 := st.ReadString()
	checkErr(err1)

	passwd, err2 := st.ReadString()
	checkErr(err2)

	if account == passwd {
		return nil
	}

	return nil

}

func checkErr(err error) {
	if err != nil {
		panic("error occured in protocol module")
	}
}
