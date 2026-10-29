package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Ip string

	Port int

	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建一个sever的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播channel的gorouline,一旦有消息就发送给全部的在线User
func (this *Server) ListrnMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部的在线User
		this.mapLock.Lock()

		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 创建一个server的接口
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg

	fmt.Println(sendMsg)

}

func (this *Server) Handler(conn net.Conn) {
	//........当前连接的业务
	user := NewUser(conn, this)
	user.Online()
	fmt.Println(user.Addr, "连接建立成功")

	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//提取用户的消息(去除'\n)
			msg := strings.TrimRight(string(buf[:n]), "\r\n")

			//用户针对msg进行消息处理
			user.DoMessage(msg)
			fmt.Println("收到", user.Addr, ":", msg)

			//用户的任意消息,代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	//当前handler堵塞
	for {
		select {
		case <-isLive:
			//当前用户是活跃的,应该重置定时器
			//不做任何事情,为了激活select,更新下面的定时器

		case <-time.After(time.Second * 20):
			//已经超时
			//将当前User强制关闭

			user.sendMsg("你被剔除了")

			//销毁用的资源
			close(user.C)

			//关闭连接
			conn.Close()

			//退出当前的Handler
			return
		}
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	//socket Listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listrn err:", err)
		return
	}

	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListrnMessager()

	for {
		//accpect

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}

}
