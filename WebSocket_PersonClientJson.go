/*
WebSocket_PersonClientJson
*/
package main

import (
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
		fmt.Println("Usage: ", os.Args[0], " ws://host:port")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "ws://localhost")
	checkError(err)

	person := Person{
		Name: "Tianle Chu",
		Emails: []string{
			"work@gmail.com",
			"life@gmail.com",
		},
	}

	err = websocket.JSON.Send(conn, person)
	if err != nil {
		fmt.Println("Couldn't send msg " + err.Error())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
