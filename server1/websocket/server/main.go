package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(res http.ResponseWriter, req *http.Request) {
	conn, err := UP.Upgrade(res, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是:%s吗?", string(p))))
		fmt.Println(t, string(p))
	}
	defer conn.Close()
	fmt.Println("服务器关闭")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
