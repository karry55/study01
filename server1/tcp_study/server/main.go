package main

import (
	"fmt"
	"net"
)

func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")

	listen, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String() + "进来了")

		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(conn.RemoteAddr().String() + "出去了")
				break
			}
			fmt.Println(string(buf[0:n]))
		}
	}
}
