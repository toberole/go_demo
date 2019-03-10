package main

import (
	"domain"
	"fmt"
	"log"
	"os"
)

// 声明结构体
type Cer struct {
	t string
	d string
}

// 声明接口
type Cert_interface interface {
	SysHello(age int, name string) (int, string)
}

func init() {
	log.SetOutput(os.Stdout)

	fmt.Println("hello init")
}
func main() {
	fmt.Println("hello world")

	res, name := domain.Add("xiaohong")

	fmt.Println(res)
	fmt.Println(name)
}
