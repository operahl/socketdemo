package main

import (
	"../utils"
	"net"
)

const (
	HOST = "localhost:1024"
	BEATINGINTERVAL = 5
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
	}

}


//handle the connection
func handleConnection(conn net.Conn, timeout int ) {

	tmpBuffer := make([]byte, 0)

	buffer := make([]byte, 1024)
	messnager := make(chan byte)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			utils.LogErr(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		tmpBuffer = utils.Depack(append(tmpBuffer, buffer[:n]...))
		utils.Log( "receive data string:", string(tmpBuffer))
		utils.TaskDeliver(tmpBuffer,conn)
		//start heartbeating
		go utils.HeartBeating(conn,messnager,timeout)
		//check if get message from client
		go utils.GravelChannel(tmpBuffer,messnager)

	}
	defer conn.Close()



}




