package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"os"
)

// 响应工具                //请求
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		data, err := os.ReadFile("hello.txt")
		if err != nil {
			fmt.Println(err)
		}
		res.Write(data)

	case "POST":
		//获取body数据
		data, err := io.ReadAll(req.Body)
		//拿请求头
		contentType := req.Header.Get("Content-Type")
		fmt.Println(contentType)

		if err != nil {
			fmt.Println(err)
		}

		ma := make(map[string]string)
		json.Unmarshal(data, &ma)
		fmt.Println(ma["username"])

		type User struct {
			Username string `json:"username"`
		}
		var user User
		json.Unmarshal(data, &user)
		fmt.Println(user)

		//设置响应头
		header := res.Header()
		header["token"] = []string{"y1gyf156sdgT%d44hjgj"}
		res.Write([]byte("hello 凤凤 POST"))
	}
}

func main() {
	//绑定回调
	http.HandleFunc("/index", IndexHandler)
	//注册服务
	fmt.Println("listen server on http://localhost:8080")
	// 改动：接住错误，端口被占用时不再静默退出
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("服务启动失败:", err)
	}
}
