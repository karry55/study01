package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Order struct {
// 	Name  string
// 	Money int
// }

// var orderChan = make(chan Order)

// func send(name string, money int, wait *sync.WaitGroup) {
// 	defer wait.Done()
// 	fmt.Printf("%s 开始购物\n", name)
// 	time.Sleep(time.Second)
// 	fmt.Printf("%s 购物结束\n", name)
// 	orderChan <- Order{Name: name, Money: money}
// }

// func main() {
// 	var wait sync.WaitGroup
// 	start := time.Now()
// 	wait.Add(3)
// 	go send("lkl", 2, &wait)
// 	go send("clx", 8, &wait)
// 	go send("lyy", 3, &wait)

// 	go func() {
// 		wait.Wait()
// 		close(orderChan)
// 	}()

// 	var moneyList []int
// 	var nameList []string
// 	for order := range orderChan {
// 		moneyList = append(moneyList, order.Money)
// 		nameList = append(nameList, order.Name)
// 	}

// 	fmt.Println("购买完成", time.Since(start))
// 	fmt.Println("moneyList:", moneyList)
// 	fmt.Println("nameList:", nameList)
// }
