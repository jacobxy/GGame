package server

import (
	"errors"
	"net"
	. "object"
	. "stream"
)

const (
	OUT_MAX = 100
)

type ConnWrite struct {
	conn net.Conn
	out  chan []byte
	ctrl chan bool
	pl   *Player
}

func NewConnWrite(con net.Conn) *ConnWrite {
	wt := &ConnWrite{conn: con, out: make(chan []byte, OUT_MAX), ctrl: make(chan bool), pl: nil}
	return wt
}

func (wt *ConnWrite) Start() {
	defer func() {
		if x := recover(); x != nil {
			//error("error on write")
		}
	}()

	for {
		select {
		case data := <-wt.out:
			wt.sendBytes(data)
		case <-wt.ctrl:
			close(wt.out)
			for data := range wt.out {
				wt.sendBytes(data)
			}
			wt.conn.Close()
		}
	}
}

func (wt *ConnWrite) sendBytes(data []byte) {
	st := GetStream()
	st.WriteBytes(data)

	_, err := wt.conn.Write(st.Data())
	if err != nil {
		errors.New("error on write")
	}
	return
}
