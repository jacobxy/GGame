package object

import (
	"db"
	"encoding/json"
	"fmt"
)

type Clan struct {
	Id            uint32
	Name          string
	PicIndex      uint8
	Announcement  string
	Announcement2 string
	Creater       uint64
	Leader        uint64
	Level         uint32
	Contribute    uint32
	PersonMax     uint8
	vecPlayer     []*Player
}

func (this *Clan) InsertPlayer(pl *Player) {
	this.vecPlayer = append(this.vecPlayer, pl)
}

var _globalClan map[uint32]*Clan

func GetGlobalClan() map[uint32]*Clan {
	if _globalClan == nil {
		_globalClan = make(map[uint32]*Clan, 1000)
	}
	return _globalClan
}

func LoadClan() {
	rows := db.SelectFromDB("select `clanId`,`name`,`picIndex`,`announcement`,`announcement2`,`creater`,`leader`,`level`,`contribute`,`personMax`from clan")
	globalClan := GetGlobalClan()
	for rows.Next() {
		clan := Clan{vecPlayer: make([]*Player, 30)}
		err := rows.Scan(&clan.Id, &clan.Name, &clan.PicIndex, &clan.Announcement, &clan.Announcement2, &clan.Creater, &clan.Leader, &clan.Level, &clan.Contribute, &clan.PersonMax)
		checkError(err)
		globalClan[clan.Id] = &clan
		fmt.Println("Clan")
		b, _ := json.Marshal(clan)
		fmt.Println(string(b))
	}
}

type ClanPlayer struct {
	ClanId     uint32
	PlayerId   uint64
	Position   uint8
	Contribute uint32
	Entertime  uint32
}

func LoadClanPlayer() {
	rows := db.SelectFromDB("select `clanId`,`playerId`,`position`,`contribute`,`entertime` from clan_player")
	globalPlayer := GetGlobalPlayers()
	globalClan := GetGlobalClan()
	for rows.Next() {
		clanPlayer := ClanPlayer{}
		err := rows.Scan(&clanPlayer.ClanId, &clanPlayer.PlayerId, &clanPlayer.Position, &clanPlayer.Contribute, &clanPlayer.Entertime)
		checkError(err)
		clan := globalClan[clanPlayer.ClanId]
		pl := globalPlayer[clanPlayer.PlayerId]
		clan.InsertPlayer(pl)
		pl.SetClan(clan)
		//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		b, _ := json.Marshal(clanPlayer)
		fmt.Println(string(b))
	}
}

/*
func checkError(err error) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
*/
