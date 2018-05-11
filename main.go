package main

import (
	"net"
)


//port : 22522

func main() {
	/*if index, err := index.New("1"); err == nil {
		fmt.Println(index.Map["name"])
	}*/
	if conn, err := net.Dial("tcp", ":22522"); err == nil {
		conn.Write([]byte("get id=1\n"))
		select {
		}
	}
}