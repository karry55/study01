package main

// import (
// 	"fmt"
// 	"runtime/debug"
// )

// func read() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err)
// 			s := string(debug.Stack())
// 			fmt.Println(s)
// 		}
// 	}()
// 	var list = []int{2, 3}
// 	fmt.Println(list[1])
// }

// func main() {
// 	read()
// }
