package main

import (
	"../utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	SERVER = "localhost:9988"
)

type Msg struct {
	Meta   map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}


func send(conn net.Conn) {

	var headSize int
	var headBytes = make([]byte, 2)
	//接收数据
	go ReceiveData(conn)
	for i := 0; i <3; i++ {
		session:=GetSession()
		message := &Msg{
			Meta:map[string]interface{}{
			"meta":"test",
			"ID":strconv.Itoa(i),
			},
		Content: Msg{
			Meta:map[string]interface{}{
			"author":"jack",
			},
		Content:session,
		},
		}
		result,_ :=	json.Marshal(message)
		headSize =len(result)
		utils.LogErr(headSize)
		binary.BigEndian.PutUint16(headBytes, uint16(headSize))

		byteData:=utils.Packet([]byte(result))
		utils.Log("send data",byteData)
		conn.Write(byteData)

		time.Sleep(1 * time.Second)
	}

	//测试服务端超时
	time.Sleep(10 * time.Second)
	utils.Log("send over")
	defer conn.Close()
}

func ReceiveData(conn net.Conn)  {
	//buffer := make([]byte, 1024)
	//n, err := conn.Read(buffer)
	//if err != nil {
	//	utils.LogErr(conn.RemoteAddr().String(), " connection error: ", err)
	//	return
	//}
	//utils.Log( "receive data string:", string(buffer[:n]))

	//TODO 后期也要处理粘包与半包，处理规则与server端一致
	utils.Log( "receive data ...")
}

func GetSession() string{
	gs1:=time.Now().Unix()
	gs2:=strconv.FormatInt(gs1,10)
	return gs2
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", SERVER)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}


	fmt.Println("connect success")
	send(conn)


}

