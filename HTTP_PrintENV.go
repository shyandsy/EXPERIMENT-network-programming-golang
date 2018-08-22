/*
HTTP_PrintENV

显示环境变量列表
http://localhost:8000/cgi-bin/printenv

显示文件列表
http://localhost:8000/

启动
go run HTTP_PrintENV.go
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)

	http.HandleFunc("/cgi-bin/printenv", printEnv)

	err := http.ListenAndServe(":8000", nil)
	checkError(err)
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>"))
	writer.Write([]byte("\n"))
	writer.Write([]byte("<pre>"))
	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
