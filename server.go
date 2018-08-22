package main

import (
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})

		// 创建case事件
		so.On("create_case", func(data string) {
			var stb interface{}
			err = json.Unmarshal([]byte(data), &stb)
			if err != nil {
				fmt.Println("create case: 解析json失败")
			} else {
				obj := stb.(map[string]interface{})
				fmt.Println("create case: ", obj["id"])
			}
		})
	})
	server.On("create_case", func() {
		log.Println("create case")
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	//http.Handle("/socket.io/", server)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		log.Println("origin", origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})
	//http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
