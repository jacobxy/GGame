package battle

import (
	"fmt"
	"object"
)

const (
	E_attr_hp = iota
	E_attr_attack
	E_attr_defend
	E_attr_critical
	E_attr_antiKnock
	E_attr_speed
	E_attr_evade
	E_attr_hit
	E_attr_max
)

type BattleAttr struct {
	attr    [E_attr_max]uint32
	skillId []uint32
}

type BattleFighter struct {
	fgt        *Fighter
	battleAttr BattleAttr
}

func NewBattleFighter(f *Fighter) *BattleFighter {
	bf := &BattleFighter{fgt: f}
	return bf
}
