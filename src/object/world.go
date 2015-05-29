package object

import (
	//	"container"
	"time"
)

type World struct {
	Id uint32
}

var _world *World

var _sortForSorce *SliceCount

func GetWorldInstance() *World {
	if _world == nil {
		_world = &World{Id: 0}
	}
	return _world
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
