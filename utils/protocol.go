
package utils

import (

	"bytes"

	"encoding/binary"

)



const (

	ConstHeader         = "version:1"

	ConstHeaderLength   = 9

	ConstSaveDataLength = 4

)



//封包

func Packet(message []byte) []byte {

	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)

}



//解包

func Unpack(buffer []byte, readerChannel chan []byte) []byte {

	length := len(buffer)



	var i int

	for i = 0; i < length; i = i + 1 {

		//读取的内容长度少于头加存储数据字段的长度，退出继续读取
		if length < i+ConstHeaderLength+ConstSaveDataLength {

			break

		}
		//头内容是否为约定的内容
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {

			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])

			//如果都长度少于数据总长度，继续读取
			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {

				break

			}

			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]

			readerChannel <- data


			//处理粘包情况
			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1

		}

	}



	if i == length {

		return make([]byte, 0)

	}

	return buffer[i:]

}



//整形转换成字节

func IntToBytes(n int) []byte {

	x := int32(n)



	bytesBuffer := bytes.NewBuffer([]byte{})

	binary.Write(bytesBuffer, binary.BigEndian, x)

	return bytesBuffer.Bytes()

}



//字节转换成整形

func BytesToInt(b []byte) int {

	bytesBuffer := bytes.NewBuffer(b)



	var x int32

	binary.Read(bytesBuffer, binary.BigEndian, &x)



	return int(x)

}