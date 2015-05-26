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
	server.StartServer()

	//slice := object.NewSliceCount(int(0))
	//slice.Add(10)
	//slice.Add(200)
	//slice.Add(30)
	//slice.Add(40)
	//sort.Sort(slice)
	//slice.String()
}
