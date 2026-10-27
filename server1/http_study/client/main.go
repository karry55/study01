package main

import (
  "fmt"
  "io"
  "net/http"
)


func main(){
	client:=new(http.Client)

	req,_:=http.NewRequest("GET", "http://localhost:8080/index", nil)

	//发请求
	// 改动：服务端没启动时返回错误，不再因 res 为 nil 而 panic
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}

	// 改动：读取响应的错误也接住
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	fmt.Println(string(b))

}
