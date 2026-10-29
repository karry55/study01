package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	Server     string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int //当前client的模式
}

// 统一用同一个 Reader 读标准输入：
// fmt.Scanln 和 bufio 混用会丢缓冲区里的数据，而且 Scanln 遇到空格就截断，
// 消息里的空格会被吃掉。
var stdin = bufio.NewReader(os.Stdin)

// 读取一行输入，去掉行尾的 \r\n；读完（Ctrl+Z / 管道结束）时返回 "exit"
func readLine() string {
	line, err := stdin.ReadString('\n')
	if err != nil && line == "" {
		return "exit"
	}
	return strings.TrimSpace(line)
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		Server:     serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	//连接server
	conn, err := net.Dial("tcp", net.JoinHostPort(serverIp, strconv.Itoa(serverPort)))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.conn = conn

	return client
}

func (client *Client) DealResponse() {
	//一旦client.conn有数据,就直接copy到stdout标准输出上,永久堵塞监听
	io.Copy(os.Stdout, client.conn)
	fmt.Println("与服务端断开连接")
	os.Exit(0)
}

func (client *Client) menu() bool {
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Print("请选择: ")

	input := readLine()
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > 3 {
		fmt.Println(">>>请输入合法范围内的数字<<<<<")
		return false
	}

	client.flag = choice
	return true
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err: ", err)
		return
	}
}

// 修改：新增 chatOnce，把 PublicChat / PrivateChat 里重复的
// "读一行 → 判 exit → 判空 → 发送" 逻辑抽到这里。
// prefix 是消息前缀（公聊为空，私聊为 "to|张三|"），返回 false 表示退出当前聊天。
func (client *Client) chatOnce(prefix, prompt string) bool {
	fmt.Print(prompt)
	chatMsg := readLine()
	if chatMsg == "exit" {
		return false
	}
	if chatMsg == "" {
		return true
	}

	if _, err := client.conn.Write([]byte(prefix + chatMsg + "\n")); err != nil {
		fmt.Println("conn Write err: ", err)
		return false
	}
	return true
}

// 私聊模式
func (client *Client) PrivateChat() {
	client.SelectUsers()
	fmt.Print(">>>>>>>请输入聊天对象[用户名],exit退出: ")

	remoteName := readLine()
	if remoteName == "exit" || remoteName == "" {
		return
	}

	// 修改：原来这里内层 for 不读输入，会死循环空转；改用 chatOnce
	for client.chatOnce("to|"+remoteName+"|", ">>>>>>>>请输入消息内容,exit退出: ") {
	}
}

// 公聊模式
func (client *Client) PublicChat() {
	// 修改：改用 chatOnce，去掉和 PrivateChat 重复的循环体
	for client.chatOnce("", ">>>>>>>请输入聊天内容,exit退出: ") {
	}
	fmt.Println()
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名:")
	name := readLine()
	if name == "" || name == "exit" {
		return false
	}

	sendMsg := "rename|" + name + "\n"

	_, err := client.conn.Write([]byte(sendMsg))

	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	client.Name = name
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			//公聊模式
			client.PublicChat()

		case 2:
			//私聊模式
			client.PrivateChat()

		case 3:
			//更新用户名
			client.UpdateName()
		}
	}
	fmt.Println("退出客户端")
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8080
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8080, "设置服务器端口(默认是8080)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>链接服务器失败....")
		return
	}

	//单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>链接服务器成功.....")

	//启动客户端的业务
	client.Run()
}
