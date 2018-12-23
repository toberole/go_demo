package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	println(r.Body)
	io.WriteString(w, "Hello world!")
}

func main() {

	fmt.Printf("hello 你好世界")
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	http.ListenAndServe(":8000", mux)
}