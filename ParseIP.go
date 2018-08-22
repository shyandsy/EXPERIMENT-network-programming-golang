/*
ip地址解析试验

运行：
go run ParseIP.go 127.0.0.1
go run ParseIP.go 0:0:0:0:0:0:0:1
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-address\n", os.Args[0])
		os.Exit(1)
	}

	// 读取参数1
	name := os.Args[1]

	// 解析地址
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	os.Exit(0)
}
