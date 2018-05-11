package index

import (
	"net"
	"strings"
	"strconv"
	"fmt"
)

type Index struct {
	Map map[string]interface{}
}

var (
	defaultPort = 22522
	connection net.Conn
)

func CreateConnection(host string, port int) bool {
	var err error
	connection, err = net.Dial("tcp", strings.Join([]string{host, strconv.Itoa(port)}, ":"))
	if err != nil {
		return false
	}
	return true
}

func GetIndex(indexId string) (*Index) {
	if connection == nil {
		if !CreateConnection("localhost", defaultPort) {
			return nil
		}
	}

	connection.Write([]byte("get id=" + indexId + "\n"))
	buff := make([]byte, 1024)
	n, _ := connection.Read(buff)
	fmt.Println(buff[:n])

	return nil
}
