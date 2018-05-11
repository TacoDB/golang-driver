package index

import "net"

type Index struct {
	Map map[string]interface{}
}

func New(objId string) (*Index, error) {
	conn, err := net.Dial("tcp", "22522")
}

func (i *Index) Get() {

}
