package main

import (
	"../utils"
	"net"
)

const (
	HOST = "localhost:9988"//服务启动监听端口
	BEATINGINTERVAL = 5//心跳时间
)

func main() {
	startServer()
}


func startServer(){
	netListen, err := net.Listen("tcp", HOST)
	utils.CheckError(err)
	defer netListen.Close()
	utils.Log("Waiting for clients")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		utils.Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn, BEATINGINTERVAL)
	}

}


//handle the connection
func handleConnection(conn net.Conn, timeout int ) {

	tmpBuffer := make([]byte, 0)

	buffer := make([]byte, 1024)
	readerChannel:=make(chan []byte,16)
	heartbeatchan := make(chan int)

	go reader(conn,heartbeatchan,readerChannel)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			utils.LogErr(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		go utils.HeartBeating(conn,heartbeatchan,timeout)
		tmpBuffer = utils.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)

	}
	defer conn.Close()
}

func reader(conn net.Conn,heartbeatchan chan int,readerChannel chan []byte) {

	for {

		select {

		case data := <-readerChannel:
			heartbeatchan<-1//心跳
			utils.RespData(data,conn)
		}

	}

}
