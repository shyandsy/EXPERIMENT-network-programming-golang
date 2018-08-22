/*
WebSocket_PersonServerJson

本代码旨在演示如何使用websocket向server发送Person数据结构

启动服务器：
go run WebSocket_PersonServerJson.go

运行：
go run WebSocket_PersonClientJson.go ws://localhost:20000
*/
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

func ReceivedPerson(ws *websocket.Conn) {
	var person Person
	err := websocket.JSON.Receive(ws, &person)
	if err != nil {
		fmt.Println("can't receive")
	} else {
		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email : " + e)
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(ReceivedPerson))
	err := http.ListenAndServe(":20000", nil)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
