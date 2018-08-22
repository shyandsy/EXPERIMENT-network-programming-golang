/*
HTTP_ClientGet
自己构建http header并发送请求

运行：
go run HTTP_ClientGet.go http://www.google.com
*/
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " host:port")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	checkError(err)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	// 只接受utf8编码
	request.Header.Add("Acceept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	checkError(err)

	response, err := client.Do(request)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Println("got the charset ", chSet)
	if strings.ToLower(chSet) != "utf-8" {
		fmt.Println("cannot handle ", chSet)
		os.Exit(3)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		fmt.Print(string(buf[0:n]))
		if err != nil {
			os.Exit(0)
		}
	}
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	contentType = strings.ToLower(contentType)
	if contentType == "" {
		return "utf-8"
	}

	idx := strings.Index(contentType, "charset: ")
	if idx == -1 {
		return "utf-8"
	}

	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
