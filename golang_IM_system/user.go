package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	//启动监听
	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {
	//用户上线,将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息

	this.server.BroadCast(this, "已上线")
}

// 用户的下线业务
func (this *User) Offline() {
	//用户下线,将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this, "下线")
}

// 给当前user对应的客户发送消息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户都有哪些

		// 只读操作，用读锁即可
		this.server.mapLock.RLock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + "在线...\n"
			this.sendMsg(onlineMsg)
		}
		this.server.mapLock.RUnlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式:rename|张三
		newName := strings.Split(msg, "|")[1]

		//判断name是否存在（读 map 也要加锁，否则可能并发读写崩溃）
		this.server.mapLock.RLock()
		_, ok := this.server.OnlineMap[newName]
		this.server.mapLock.RUnlock()
		if ok {
			this.sendMsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			// this.sendMsg("您已经更新用户名:" + this.Name + "\n")
			fmt.Println("用户", this.Addr, "改名为", newName)
		}
	} else if strings.HasPrefix(msg, "to|") {
		//消息格式: to|张三|消息内容

		// 先校验格式再取下标，否则 "to|张三" 这种缺少内容的消息会下标越界 panic
		parts := strings.SplitN(msg, "|", 3)
		if len(parts) < 3 || parts[1] == "" || parts[2] == "" {
			this.sendMsg("消息格式不正确:请使用\"to|张三|你好啊\"格式,\n")
			return
		}

		// 1 获取对方的用户名和消息内容
		remoteName := parts[1]
		content := parts[2]

		// 2根据用户名 得到对方User对象（读 map 加读锁）
		this.server.mapLock.RLock()
		remoteUser, ok := this.server.OnlineMap[remoteName]
		this.server.mapLock.RUnlock()
		if !ok {
			this.sendMsg("该用户名不存在\n")
			return
		}

		//获取消息内容,通过对方的User对象将消息内容发送过去
		remoteUser.sendMsg(this.Name + "对您说:" + content + "\n")
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 消息广播处理
func (this *User) BroadCast(user *User, msg string) {
	this.server.BroadCast(user, msg)
}

// 监听当前User channel的方法,一旦有消息,就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
