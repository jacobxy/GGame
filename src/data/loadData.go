package data

import (
	"fmt"
)

func Load() {
	LoadDataFighter()
	LoadDataItem()
	LoadDataExp()
}

func checkError(err error) {
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
}
