/*
WebSocket_PersonClientXML

启动server
go run WebSocket_PersonServerXML

运行：
go run WebSocket_PersonClientXML.go ws://localhost:20000
*/
package main

import (
	"./xmlcodec"
	"fmt"
	"golang.org/x/net/websocket"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)

	person := Person{
		Name: "Tianle Chu",
		Emails: []string{
			"work@gmail.com",
			"life@gmail.com",
		},
	}

	// 发送
	err = xmlcodec.XMLCodec.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg: " + err.Error())
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
