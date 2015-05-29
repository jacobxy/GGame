package data

import (
	"db"
	"fmt"
)

type ItemDB struct {
	Id          uint32
	Name        string
	SubClass    uint8
	MaxQuantity uint32
	Sale        uint32
	//AttrId   uint32
}

var _globalItem map[uint32]*ItemDB

func GetGlobalItem() map[uint32]*ItemDB {
	if _globalItem == nil {
		_globalItem = make(map[uint32]*ItemDB, 1000)
	}
	return _globalItem
}

func LoadDataItem() {
	//rows := db.SelectFromData("select `id`,`name`,`subclass`,`career`,`attrId`,`coin` from item_template")
	rows := db.SelectFromData("select `id`,`name`,`subclass`,`maxQuantity`,`coin` from item_template2")
	for rows.Next() {
		item := &ItemDB{}
		err := rows.Scan(&item.Id, &item.Name, &item.SubClass, &item.MaxQuantity /* &item.AttrId,*/, &item.Sale)
		checkErr(err)
		GetGlobalItem()[item.Id] = item
	}
	fmt.Println("ItemLen:", len(GetGlobalItem()))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
