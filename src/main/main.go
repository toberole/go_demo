package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func test_hello1(w http.ResponseWriter) {
	fmt.Println("test_hello1")
	io.WriteString(w, "test_hello1 ")
}

func test_hello2(ch chan string) {
	fmt.Println("test_hello2")
	time.Sleep(time.Duration(2) * time.Second)
	ch <- "test_hello2 Sleep 2 Second"
}

// 按照规定的签名定义 作为http的处理函数
func hello1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world! ")
	test_hello1(w)

	// 注意启动的线程必须在服务回调的方法之前结束
	// 否则 http请求结束了 线程再返回结果就无意义
	// 利用chan阻塞特性
	var ch = make(chan string)
	go test_hello2(ch)
	var v = <-ch
	io.WriteString(w, v)
}

// 获取get 参数
func get_param1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("===========get_param==========")

	query := r.URL.Query()
	name := query["name"][0]
	fmt.Println(name)

	w.Write([]byte("hello get"))
}

// 获取post方法的不带文件的参数
func post_param1(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	pwd := r.PostFormValue("pwd")

	fmt.Printf("name: %s, pwd: %s", name, pwd)

	w.Write([]byte(" hello post_param"))
}

func main() {
	fmt.Println("hello 你好世界")

	// 路由 可以使用DefaultServeMux 默认的
	mux := http.NewServeMux()

	// 注册一个处理函数
	mux.HandleFunc("/hello", hello1)
	mux.HandleFunc("/get_param", get_param1)
	mux.HandleFunc("/post_param", post_param1)

	// 启动服务 监听8000端口
	http.ListenAndServe(":8000", mux)
}
