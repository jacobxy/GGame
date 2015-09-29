package server

import (
	//"net"
	"object"
	. "stream"
)

var NetHandler map[uint16]func(*ConnWrite, *Stream) interface{}

func init() {
	NetHandler = map[uint16]func(*ConnWrite, *Stream) interface{}{
		1: handleLogin,
		2: LoadPlayerPtr,
	}
}

func checkErr(err error) {
	if err != nil {
		panic("error occured in protocol module")
	}
}

func handleLogin(wt *ConnWrite, st *Stream) interface{} {

	account, err1 := st.ReadString()
	checkErr(err1)

	passwd, err2 := st.ReadString()
	checkErr(err2)

	if account == passwd {
		return nil
	}
	return nil

	playerId, err := object.LoginForPid(account, passwd)
	checkErr(err)
	return playerId
}

func LoadPlayerPtr(out *ConnWrite, st *Stream) interface{} {
	playerId, err := st.ReadU64()
	checkErr(err)
	pl, ok := object.GetGlobalPlayers()[playerId]
	if ok {
		out.pl = pl
	} else {
		out.conn.Write(([]byte)("error No Player"))
		return nil
	}
	return pl.Name
}
