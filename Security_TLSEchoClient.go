/*
Security_TLSEchoClient

本代码旨在测试带有TLS(传输层加密)功能的echo server，Security_TLSEchoClient。

原始代码中，tls.dial没有使用参数InsecureSkipVerify: true配置，导致系统报错
unknown authority
*/
package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " host:port")
		os.Exit(1)
	}

	service := os.Args[1]

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", service, config)
	checkError(err)

	for n := 0; n < 10; n++ {
		fmt.Println("Writing....")
		conn.Write([]byte("hello world!! " + string(n+48)))

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
