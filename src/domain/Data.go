package domain

import "fmt"

// 数组
func Test1() {
	// 声明一个数组 在Go语言中 声明变量的时候 总会使用对应类型的零值初始化
	// var array [5]int
	// array := [5]int{1,2,3,4,5}
	array := [...]int{1, 2, 3, 4, 5}

	// 数组也是个变量 赋值操作会导致数组里面的所有值拷贝一份到被赋值的数组
	array1 := array

	// 使用数组
	array[2] = 110
	fmt.Println(array[2])
	fmt.Println(array1[2])

	// 多维数组 两行四列的数组
	// arrs:=[4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	// 声明数组 并且初始化第一个和第三个元素
	arrs := [4][2]int{1: {111, 222}, 3: {333, 444}}

	fmt.Println(arrs[1][0])
}

// 切片
func Test2() {
	// make 创建切片
	// slice := make([]string,5)
	// 创建一个长度为5，容量为10的切片 容量必须>=长度
	// slice := make([]string, 5, 10)

	// 使用类似创建数组的方法创建切片 长度为2，容量为2
	// slice := []string{"xiaohon", "xiaoming"}
	// 长度为100，容量为100
	// slice := []string{99: ""}

	// 创建nil整形切片
	// var slice []int
	// 使用make创建nil整形切片
	// slice := make([]int,0)
	// 使用切片字面量创建空的整形切片
	// slice := []int{}

	//slice := []int{1, 2, 3, 4}
	//slice[0] = 2222
	//fmt.Println(slice[0])

	// 使用切片创建切片 原始切片和新创建的切片共享底层的数组
	//slice := []int{1, 2, 3, 4, 5}
	//// 使用slice创建一个长度为2，容量为4的切片
	//newSlice := slice[1:3]// 包头不包尾 底层共享 slice[1],slice[2]
	//
	//newSlice[0] = 9
	//fmt.Println(newSlice[0])
	//fmt.Println(slice[1])

	/**
	 	对于底层数组容量是K的切片 slice[i:j]
		长度：j - i
		容量：k - i
	*/

	//slice := []int{1, 2, 3, 4, 5}
	//slice1 := slice[1:3]
	//append(slice1, 10) // 此时slice1和slice依旧共享底层数组 slice[3]==10

	//slice := []int{1, 2, 3, 4, 5}
	//slice1 := append(slice, 100) // 此时系统会创建一个新的底层数组，并且容量是以前的两倍

	/**
	 对于底层数组容量是K的切片 slice[i:j:k]
		长度：j - i
		容量：k - i
	i: 表示新切片开始的索引位置
	j：表示 i+新切片包含元素的个数
	k: 表示 i+新切片容量
	新切片的容量不允许大于旧切片的容量
	*/

	slice := []int{5, 4, 3, 2, 1}
	// for range 迭代切片
	for index, value := range slice {
		fmt.Printf("index: %d value: %d\n", index, value)
	}

	fmt.Println("------------------------------------")

	for index := 2; index < len(slice); index++ {
		fmt.Printf("index: %d value: %d\n", index, slice[index])
	}
}

// map
func Test3() {
	// 创建一个key为string，value为int
	// dict:=make(map[string]int)
	dict := map[string]int{"1": 1, "2": 2}
	// 获取key为"1"的值 同时判断存在与否
	value, exists := dict["1"]
	if exists {
		fmt.Println(value)
	}

	for key, value := range dict {
		fmt.Printf("key: %s value: %d\n", key, value)
	}

	// 删除
	delete(dict, "1")

}
