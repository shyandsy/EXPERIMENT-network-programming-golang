/*
HTTP_Get

本代码用于获取网站http response的header
原始代码中判断utf8编码的方法有缺陷，因为没有处理大小写问题。html中文档编码并不限制大写或者小写。

运行
go run HTTP_Get.go http://www.golang.com/
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " host:port")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(3)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))

	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		fmt.Print(string(buf[0:n]))
		if err != nil {
			os.Exit(0)
		}
	}
	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(strings.ToLower(cType), "utf-8") != -1 {
			return true
		}
	}
	return false
}
