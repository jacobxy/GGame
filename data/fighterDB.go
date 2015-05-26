package data

import (
	"db"
	"encoding/json"
	"fmt"
)

type Fighter struct {
	Id          uint32
	Name        string
	Color       uint8
	TypeId      uint8
	ChildType   uint8
	Speed       uint32
	BodySize    uint8
	Skills      string
	Hp          uint32
	Attack      uint32
	Defend      uint32
	MagAtk      uint32
	MagDef      uint32
	Critical    uint32
	CriticalDef uint32
	Hit         uint32
	Evade       uint32
}

var _globalFighters map[uint32]*Fighter

func GetGlobalFighters() map[uint32]*Fighter {
	if _globalFighters == nil {
		_globalFighters = make(map[uint32]*Fighter)
	}
	return _globalFighters
}

func (fighter *Fighter) Clone() *Fighter {
	fgt := Fighter{}

	fgt.Id = fighter.Id
	fgt.Name = fighter.Name
	fgt.Color = fighter.Color
	fgt.TypeId = fighter.TypeId
	fgt.ChildType = fighter.ChildType
	fgt.Speed = fighter.Speed
	fgt.BodySize = fighter.BodySize
	fgt.Skills = fighter.Skills
	fgt.Hp = fighter.Hp
	fgt.Attack = fighter.Attack
	fgt.Defend = fighter.Defend
	fgt.MagAtk = fighter.MagAtk
	fgt.MagDef = fighter.MagDef
	fgt.Critical = fighter.Critical
	fgt.CriticalDef = fighter.CriticalDef
	fgt.Hit = fighter.Hit
	fgt.Evade = fighter.Evade

	return &fgt
}

func LoadDataFighter() {
	rows := db.SelectFromData("select * from fighter_base")
	globalFighter := GetGlobalFighters()

	for rows.Next() {
		var id uint32
		fighter := Fighter{}
		err := rows.Scan(&id, &fighter.Name, &fighter.Color, &fighter.TypeId, &fighter.ChildType, &fighter.Speed,
			&fighter.BodySize, &fighter.Skills, &fighter.Hp, &fighter.Attack, &fighter.Defend, &fighter.MagAtk, &fighter.MagDef, &fighter.Critical, &fighter.CriticalDef, &fighter.Hit, &fighter.Evade)
		globalFighter[id] = &fighter
		checkError(err)

		b, _ := json.Marshal(fighter)
		fmt.Println(string(b))
	}
}
