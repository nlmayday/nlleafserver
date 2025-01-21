package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
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

func TestWebsocket(t *testing.T) {

	data := []byte(`{"hdr":{"subid":10001,"act":"keepalive"},"body":{"time":1}}`)
	// hdrMap := map[string]interface{}{"subid": 10001, "act": "keepalive"}
	// bodyMap := map[string]interface{}{"time": 1}
	// dataMap := map[string]interface{}{"hdr": hdrMap, "body": bodyMap}
	// bbody, _ := json.Marshal(bodyMap)
	// dataMap["body"] = string(bbody)

	// data, _ := json.Marshal(dataMap)

	// len + data
	m := make([]byte, 10+len(data))

	// 前十个字节为空格
	// copy(m[10:], data)
	for i := 0; i < 10; i++ {
		m[i] = ' '
	}
	// 默认使用大端序
	// binary.BigEndian.PutUint16(m, uint16(78))
	lenStr := fmt.Sprintf("%d", len(data))
	copy(m[0:], lenStr)
	copy(m[10:], data)

	log.Printf("send: %v\n", m)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "server.volcanotest.com", Path: "/"} // 修改为你的 WebSocket服务器 地址
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// 发送一次消息"A"

	err = c.WriteMessage(websocket.BinaryMessage, m)
	if err != nil {
		log.Fatal("write:", err)
	}
	log.Println("sent: ", string(m))

	// 设置接收消息的数量
	n := 10 // 假设你想接收10条消息
	// 开始接收消息
	for i := 0; i < n; i++ {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %v\n", message)
		log.Printf("recv: %s", message)
	}

	// 处理中断信号
	select {
	case <-interrupt:
		log.Println("interrupt")

		// 清洁关闭连接
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
		}
		select {
		case <-time.After(time.Second):
		case <-interrupt:
		}
	}
}
