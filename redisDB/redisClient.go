package redisDB

import (
	"bufio"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"net"
)

type redisClient struct {
	redisDb    *db
	scoketName string
	name       string
	cmd        *Cmd
	bufIn      bufio.Reader
	bufOut     bufio.Writer
}

type Clients struct {
}
type Cmd struct {

}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}