package utils

import (
	"net"
	"time"
)


//HeartBeating, determine if client send a message within set time by GravelChannel

func HeartBeating(conn net.Conn, readerChannel chan byte,timeout int) {
	select {
	case _ = <-readerChannel:
		Log("get message from "+conn.RemoteAddr().String()+ ", keeping heartbeating...")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second*5):
		Log("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func GravelChannel(n []byte,mess chan byte){
	for _ , v := range n{
		mess <- v
	}
	close(mess)
}

