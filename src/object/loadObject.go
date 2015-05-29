package object

import (
	"data"
	"db"
	"encoding/json"
	"fmt"
	//"global"
)

type DBVar struct {
	PlayerId uint64 `json:"playerId"`
	Id       uint32 `json:"varId"`
	Data     uint32 `json:"data"`
	Over     uint32 `json:"over"`
}

func loadVar() {
	rows := db.SelectFromDB("select * from var ")
	for rows.Next() {
		usr := DBVar{}
		err := rows.Scan(&usr.PlayerId, &usr.Id, &usr.Data, &usr.Over)
		checkError(err)
		player := GetGlobalPlayers()[usr.PlayerId]
		if player == nil {
			fmt.Println("player == nil")
			b, _ := json.Marshal(usr)
			fmt.Println(string(b))
			continue
		}
		player.MyV.LoadVar(usr.Id, usr.Data, usr.Over)

		b, _ := json.Marshal(usr)
		fmt.Println(string(b))
	}
}

type DBPlayer struct {
	Id    uint64
	Name  string
	Level uint8
}

func loadPlayer() {
	rows := db.SelectFromDB("select id,name,34 from player")
	globalPlayer := GetGlobalPlayers()
	for rows.Next() {
		player := DBPlayer{}
		err := rows.Scan(&player.Id, &player.Name, &player.Level)
		checkError(err)
		one := NewPlayer()
		one.LoadPlayerInfo(player.Id, player.Name, player.Level)
		globalPlayer[player.Id] = one

		b, _ := json.Marshal(player)
		fmt.Println(string(b))
	}
}

type DBFighter struct {
	PlayerId   uint64
	Figther    uint32
	Experience uint64
	AddTime    uint32
}

func loadFighter() {
	var playerId uint64
	var fighterId uint32
	var exp uint32
	var addTime uint32

	rows := db.SelectFromDB("select playerId, fighterId, experience, addTime from fighter")

	for rows.Next() {
		err := rows.Scan(&playerId, &fighterId, &exp, &addTime)
		checkError(err)

		fmt.Println("XXXXXXXX", playerId, fighterId, exp, addTime)

		player := GetGlobalPlayers()[playerId]
		if player == nil {
			continue
		}

		fighters := player.GetFighters()
		if fighters == nil {
			continue
		}

		if fighters[fighterId] != nil {
			continue
		}

		fighter := data.GetGlobalFighters()[fighterId]
		if fighter == nil {
			continue
		}

		name := player.Name
		if fighterId > 10 {
			name = fighter.Name
		}

		TheFighter := &Fighter{}
		TheFighter.LoadPlayerInfo(fighterId, name, exp, addTime)

		fmt.Println(TheFighter)
		fighters[fighterId] = TheFighter
	}
}

func Load() {
	loadPlayer()
	loadFighter()
	loadVar()
	LoadClan()
	LoadClanPlayer()
}

func checkError(err error) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
