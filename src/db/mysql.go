package db

import (
	"database/sql"
	//"encoding/json"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var _mysql *sql.DB
var _mysqlData *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Instance() *sql.DB {
	var err error
	if _mysql == nil {
		_mysql, err = sql.Open("mysql", "jacob:kingxin@/asss_hand")
		checkErr(err)
	}
	return _mysql
}

func Option(str string) {
	my := Instance()
	stmt, err := my.Prepare(str)
	checkErr(err)
	defer stmt.Close()
	stmt.Exec()
}

func SelectFromDB(str string) *sql.Rows {
	my := Instance()
	rows, err := my.Query(str)
	checkErr(err)
	return rows
}

func InstanceData() *sql.DB {
	var err error
	if _mysqlData == nil {
		_mysqlData, err = sql.Open("mysql", "jacob:kingxin@/data_hand")
		checkErr(err)
	}
	return _mysqlData
}

func SelectFromData(str string) *sql.Rows {
	my := InstanceData()
	rows, err := my.Query(str)
	checkErr(err)
	return rows
}
