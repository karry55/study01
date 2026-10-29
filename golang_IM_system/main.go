package main

func main() {
	server := NewServer("127.0.0.1", 8080)
	server.Start()
}

// 程序入口：没有 main 函数时 package main 无法编译成可执行程序
func main() {
	runServer()
}
