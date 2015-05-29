package object

import (
//"object"
)

var GlobalPlayers map[uint64]*Player

func GetGlobalPlayers() map[uint64]*Player {
	if GlobalPlayers == nil {
		GlobalPlayers = make(map[uint64]*Player)
	}
	return GlobalPlayers
}
