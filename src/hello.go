package main

import (
	"fmt"
	"io"
	"net/http"
)

// 按照规定的签名定义 作为http的处理函数
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func get_param(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("===========get_param==========")

	query := r.URL.Query()
	name := query["name"][0]
	fmt.Println(name)

	w.Write([]byte("hello get"))
}

func main() {
	fmt.Println("hello 你好世界")

	// 路由 可以使用DefaultServeMux 默认的
	mux := http.NewServeMux()

	// 注册一个处理函数
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/get_param", get_param)

	// 启动服务 监听8000端口
	http.ListenAndServe(":8000", mux)
}