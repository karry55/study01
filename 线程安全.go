package main

// import (
// 	"fmt"
// 	"sync"
// )

// var num int
// var wait sync.WaitGroup
// var lock sync.Mutex

// func add() {
// 	lock.Lock()
// 	for i := 0; i < 100000; i++ {
// 		num++
// 	}
// 	lock.Unlock()
// 	wait.Done()
// }

// func reduce() {
// 	lock.Lock()
// 	for i := 0; i < 100000; i++ {
// 		num--
// 	}
// 	lock.Unlock()
// 	wait.Done()
// }

// func main() {
// 	wait.Add(2)
// 	go add()
// 	go reduce()
// 	wait.Wait()
// 	fmt.Println(num)
// }
