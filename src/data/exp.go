package data

import (
	"db"
	"fmt"
)

type ExpCount uint64

var _globalExp map[uint8]ExpCount

func GetGlobalExp() map[uint8]ExpCount {
	if _globalExp == nil {
		_globalExp = make(map[uint8]ExpCount)
	}
	return _globalExp
}

func LoadDataExp() {
	rows := db.SelectFromData("select `lvl`,`exp` from  `lvl_exp`")
	for rows.Next() {
		//item := &ItemDB{}
		var level uint8
		var exp ExpCount
		err := rows.Scan(&level, &exp)
		checkErr(err)
		GetGlobalExp()[level] = exp
	}
	fmt.Println("ExpLen:", len(GetGlobalExp()))
}
