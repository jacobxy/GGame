package object

import (
//"fmt"
//"strings"
//"data"
)

type Fighter struct {
	Id   uint32
	Name string
	Exp  uint32
	//base    *data.Fighter
	AddTime uint32
	MyV     *Var
}

func (fighter *Fighter) LoadPlayerInfo(id uint32, name string, exp uint32, addTime uint32) {
	fighter.Id = id
	fighter.Name = name
	fighter.Exp = exp
	fighter.AddTime = addTime
	//fighter.base = data.GetGlobalFighters()[id]
}
