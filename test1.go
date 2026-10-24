package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello, World!")
}

func main() {
	// for i := 0; true; i++ {
	// 	fmt.Println(time.Now())
	// 	time.Sleep(2 * time.Second)
	// }
	fmt.Println("Hello, World!")
	fmt.Printf("%v %Ts\n", "ff", "ff")
	fmt.Printf("%d\n", 3)
	var list []string

	list = append(list, "枫枫")
	list = append(list, "知道")
	fmt.Println(list)
	fmt.Println(len(list)) // 切片长度
	// 修改第二个元素
	list[1] = "不知道"
	fmt.Println(list)
	list1 := make([]int, 2, 2)
	fmt.Println(list1, len(list1), cap(list1))
}
