/*
WebSocket_EchoServer

在接收链接之后，向客户端发送数据

启动服务器：
go run WebSocket_EchoServer.go

*/
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing")

	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+48)
		fmt.Println("Sending to client: " + msg)

		// 发送数据到客户端
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("can't send message")
			break
		}

		// 接收数据
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("cant receive message")
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	err := http.ListenAndServe(":20000", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
