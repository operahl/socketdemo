package main

import (
	"../utils"
	"net"
)

const (
	HOST = "localhost:9988"
	BEATINGINTERVAL = 5
	BYTES_SIZE uint16 = 1024
	HEAD_SIZE  int    = 2
)

func main() {
	startServer("./conf/config.yaml")
}


func startServer(configpath string){
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
		//go doConn(conn)
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
