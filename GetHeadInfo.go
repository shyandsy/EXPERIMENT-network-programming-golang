/*
GetHeadInfo

go run GetHeadInfo.go www.google.com:80
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	// 解析IP地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 建立连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// 发出请求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// 读取响应
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
