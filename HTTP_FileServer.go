/*
HTTP_FileServer

启动之后浏览器访问http://localhost:8000/即可查看当前目录下文件

go run HTTP_FileServer.go
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("./"))

	err := http.ListenAndServe(":8000", fileServer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
