package main

// 修改：合并成唯一的 main（之前 runServer + main 重复导致 main redeclared）
func main() {
	server := NewServer("127.0.0.1", 8080)
	server.Start()
}
