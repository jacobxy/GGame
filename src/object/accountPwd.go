package object

import (
	"db"
	"errors"
	"fmt"
)

type accountPwd struct {
	//pid      uint64
	account  string
	password string
}

var globalAccount map[accountPwd]uint64

func init() {
	globalAccount = make(map[accountPwd]uint64)
}

func loadAccountPwd() {
	rows := db.SelectFromDB("select *from player_id")
	for rows.Next() {
		var pid uint64
		var acc string
		var pwd string
		err := rows.Scan(&pid, &acc, &pwd)
		checkError(err)
		globalAccount[accountPwd{acc, pwd}] = pid

		fmt.Println("账号:", acc, " 密码:", pwd, " 玩家ID:", pid)
	}
}

func LoginForPid(acc, pwd string) (uint64, error) {
	pid, ok := globalAccount[accountPwd{acc, pwd}]

	if !ok {
		return 0, errors.New("error")
	}

	return pid, nil
}
