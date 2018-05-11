package main

import "mvmo.eu/tacodb/database/index"

//port : 22522

func main() {
	/*if index, err := index.New("1"); err == nil {
		fmt.Println(index.Map["name"])
	}*/
	index.GetIndex("1")

}