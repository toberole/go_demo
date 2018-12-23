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

// 获取get 参数
func get_param(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===========get_param==========")

	query := r.URL.Query()
	name := query["name"][0]
	fmt.Println(name)

	w.Write([]byte("hello get"))
}

// 获取post方法的不带文件的参数
func post_param(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=============post_param==============")

	name := r.PostFormValue("name")
	pwd := r.PostFormValue("pwd")

	fmt.Printf("name: %s, pwd: %s",name,pwd)

	w.Write([]byte(" hello post_param"))
}

func main() {
	fmt.Println("hello 你好世界")

	// 路由 可以使用DefaultServeMux 默认的
	mux := http.NewServeMux()

	// 注册一个处理函数
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/get_param", get_param)
	mux.HandleFunc("/post_param", post_param)

	// 启动服务 监听8000端口
	http.ListenAndServe(":8000", mux)
}
