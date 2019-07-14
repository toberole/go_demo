package main

import (
	xxxx "domain" // 导入domain包 取别名为xxxx
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

func user_interface(i xxxx.Interface_xxxx) {
	i.Sys_xxx()
}

func main2() {
	fmt.Println("hello world")

	res, name := xxxx.Add("xiaohong")

	// xxxx.Test1()
	// xxxx.Test2()
	xxxx.Test3()

	//cert := Cer{
	//	t: "test",
	//	d: "xxx",
	//}
	//cert := Cer{
	//	"test",
	//	"xxx",
	//}
	//fmt.Println(cert.t)
	//fmt.Println(cert.d)

	user := xxxx.User{"xiaoming", 11}
	user.SysHellouser()

	user1 := &xxxx.User{"xiaohong", 22}
	user1.SysHellouser1()

	user2 := &xxxx.User{"xiaohong----", 33}
	user2.SysHellouser2()
	fmt.Println(user2.Name)

	fmt.Println("======================================")
	user3 := &xxxx.User{"xiaohong----", 55555}
	user_interface(user3)

	fmt.Println(res)
	fmt.Println(name)
}
