/*
HTTP_ServerHandler

对所有请求返回204 no content
*/
package main

import "net/http"

func main() {
	myHandler := http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusNoContent)
	})
	http.ListenAndServe(":8000", myHandler)
}
