package domain

import "fmt"

/**
Go 语言函数定义
func [方法的接受者声明,类的畸形实现 m default] func_name(参数列表)(返回值类型列表){
	// 函数体
}
*/
func Add(name string) (int32, string) {
	fmt.Println("add a student")
	fmt.Println(name)
	return 0, "xiaogang"
}
