package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("hello_world.txt")
// 	if err != nil {
// 		fmt.Println("打开文件失败:", err)
// 		return
// 	}
// 	defer file.Close()

// 	buf := bufio.NewReader(file)

// 	for {
// 		line, _, err := buf.ReadLine()

// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("读取错误:", err)
// 			break
// 		}

// 		fmt.Println(string(line))

// 	}
// }
