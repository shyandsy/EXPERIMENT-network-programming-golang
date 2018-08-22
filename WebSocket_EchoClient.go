/*
WebSocket_EchoClient

本代码演示了如何接受服务器发来的数据

启动服务器：
go run WebSocket_EchoServer.go

运行
go run WebSocket_EchoClient.go ws://localhost:20000
*/

package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)

	var msg string
	for {
		// 接收数据
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// 接受完毕，正常退出
				break
			}
			fmt.Println("Couldn't receive message: " + err.Error())
			break
		}

		fmt.Println("Received from server: " + msg)

		// 发回数据
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
