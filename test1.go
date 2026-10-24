package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, World!")
}

func main() {
	for i := 0; true; i++ {
		fmt.Println(time.Now())
		time.Sleep(2 * time.Second)
	}
}
