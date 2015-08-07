package object

import (
	//	"container"
	"stream"
	"time"
)

type Message struct {
	Pl  *Player
	Msg stream.Stream
}

type World struct {
	Id  uint32
	Mes chan Message
}

var _world *World

var _sortForSorce *SliceCount

type handlerFunc func(pl *Player, str stream.Stream) bool

var _worldFunc map[uint16]handlerFunc

func init() {
	_worldFunc = make(map[uint16]handlerFunc)
	_worldFunc[0] = Say
}

func GetWorldInstance() *World {
	if _world == nil {
		_world = &World{Id: 0, Mes: make(chan Message)}
		go HandlerMessage(_world)
	}
	return _world
}

func Say(pl *Player, context stream.Stream) bool {
	str := context.ReadString()

}

func EnterMessage(pl *Player, st stream.Stream) {
	GetWorldInstance().Mes <- Message{pl, st}
}

//handler the World Message
func HandlerMessage(wd *World) {
	for {
		select {
		case worldMsg := <-wd.Mes:
			funcId := worldMsg.Msg.ReadU16()
			fn, ok := _worldFunc[funcId]
			if ok {
				fn(worldMsg.Pl, worldMsg.Msg)
			}
		}
	}
}

func AddTime(timeCount int32, fc func() bool) {
	time2 := time.NewTicker(time.Duration(timeCount) * time.Second)
	for {
		select {
		case <-time2.C:
			fc()
		}
	}
}

func GetSortForScore() *SliceCount {
	if _sortForSorce == nil {
		_sortForSorce = NewSliceCount(uint32(0))
	}
	return _sortForSorce
}
