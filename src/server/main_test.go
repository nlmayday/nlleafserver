package main

import (
	"encoding/binary"
	"net"
	"testing"
)

// 创建一个ws连接
func TestLogin(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	data := []byte(`{
		"SGMessage":{
			"hdr" : {
				"subid" : 10200,
				"act" : "user_pwd"
			},
			"body" : {
				"dev" : "dev",
				"version" : "version",
				"descrip" : "descrip",
				"systype" : "systype",
				"type" : "type",
				"username" : "username",
				"checkinfo" : "checkinfo",
				"expinfo" : "expinfo"
			}
		}
	}
	`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)

	for {
		// 读取消息
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		println(string(buf[:n]))
	}
}
