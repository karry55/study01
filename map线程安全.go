package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wait sync.WaitGroup
// var mp = sync.Map{}

// func reader() {
// 	for {
// 		fmt.Println(mp.Load("time"))
// 	}
// 	wait.Done()
// }

// func writer() {
// 	for {
// 		mp.Store("time", time.Now().Format("15:04:05"))
// 	}
// 	wait.Done()
// }

// func main() {
// 	wait.Add(2)
// 	go reader()
// 	go writer()
// 	wait.Wait()
// }
