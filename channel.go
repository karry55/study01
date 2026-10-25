package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var moneyChan = make(chan int)

// func pay(name string, money int, wait *sync.WaitGroup) {
// 	fmt.Println("%s 开始购物", name)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("%s 购物结束\n", name)

// 	moneyChan <- money

// 	wait.Done()
// }

// func main() {
// 	var wait sync.WaitGroup
// 	startTime := time.Now()

// 	wait.Add(3)

// 	go pay("张三", 2, &wait)
// 	go pay("lkl", 4, &wait)
// 	go pay("clx", 7, &wait)

// 	go func() {
// 		defer close(moneyChan)
// 		wait.Wait()
// 	}()

// 	for {
// 		money, ok := <-moneyChan
// 		fmt.Println(money, ok)
// 		if !ok {
// 			break
// 		}
// 	}

// 	fmt.Println("购买完成", time.Since(startTime))

// }
