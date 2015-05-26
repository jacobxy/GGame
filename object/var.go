package object

import (
	"time"
)

const (
	VAR_NONE = iota
)

type Var struct {
	MyVar     map[uint32]uint32
	MyVarOver map[uint32]uint32
}

func (_var *Var) LoadVar(index uint32, value uint32, time uint32) {
	_var.MyVar[index] = value
	_var.MyVarOver[index] = time
}

func (_var *Var) GetVar(index uint32) uint32 {
	now := time.Now().Unix()
	if uint32(now) >= _var.MyVarOver[index] {
		_var.MyVar[index] = 0
	}
	return _var.MyVar[index]
}
