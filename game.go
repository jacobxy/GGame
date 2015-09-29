package main

import (
	//"db"
	//"encoding/json"
	"data"
	"fmt"
	"object"
	"server"
	//"sort"
)

func main() {
	data.Load()
	object.Load()

	for key, value := range object.GetGlobalPlayers() {
		fmt.Println(key, value)
		fighters := value.GetFighters()
		fmt.Println("---------")
		fmt.Println(fighters)
		for key1, value1 := range fighters {
			fmt.Println(key1, value1)
		}
	}
	fmt.Println("complete")
	server.StartServer2()
}
