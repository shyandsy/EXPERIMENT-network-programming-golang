/*
WebSocket_PersonServerXML

运行
go run WebSocket_PersonServerXML
*/
package main

import (
	"./xmlcodec"
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
	err := xmlcodec.XMLCodec.Receive(ws, &person)
	if err != nil {
		fmt.Println("Cant receive")
	} else {
		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
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
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
