package domain

import "fmt"

/**
Go 语言函数定义
func [接受者声明,类的畸形实现 m default] func_name(参数列表)(返回值类型列表){
	// 函数体
}

如果一个函数有接受者，那么这个函数一般称为方法
*/
func Add(name string) (int32, string) {
	fmt.Println("add a student")
	fmt.Println(name)
	return 0, "xiaogang"
}

// 声明一个结构体
type User struct {
	Name string
	Age  int
}

// 给结构体添加方法 类的畸形实现
func (u User) SysHellouser() {
	fmt.Println(u.Name)
	fmt.Println(u.Age)
}

func (u *User) SysHellouser1() {
	fmt.Println(u.Name)
	fmt.Println(u.Age)
}

func (u *User) SysHellouser2() {
	u.Name = "SysHellouser2"
	fmt.Println(u.Name)
}
