package test1

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32 = 0
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

// 共享资源引起问题
func incre() {
	defer wg.Done()
	for index := 0; index < 10; index++ {
		// counter++
		// 采用原子操作 对counter加1操作时原子性的操作
		// atomic.LoadInt32() 原子性的读
		// atomic.StoreInt32() 原子性的写
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

// 锁机制保证同步
func incre1() {
	defer wg.Done()
	for index := 0; index < 10; index++ {
		// 加锁访问counter
		mutex.Lock()
		counter++
		runtime.Gosched()
		mutex.Unlock()
	}
}

// 通道可以共享内置类型、命名类型、结构类型和引用类型的值或者指针
// 无缓冲通道
var unbuffered = make(chan int)

// 有缓冲通道
var buffered = make(chan string, 10)

/**
向通道发送数据 buffered <- "hello"
从通道读取数据 value := <- buffered
*/

func main() {

	wg.Add(3)

	go incre1()
	go incre1()
	go incre1()

	wg.Wait()

	fmt.Printf("all task has done counter = %d\n", counter)
}

// Go 语言的goroutine是由其内部的任务调度器调度运行的
// 有别于操作系统调度任务
/*
func test1() {
	// 打印cpu的核数
	fmt.Println(runtime.NumCPU())

	// 设置逻辑处理器的个数
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("runnable 1")

	}()

	go func() {
		defer wg.Done()
		fmt.Println("runnable 2")
		runtime.Gosched() // 放弃cpu执行权
		fmt.Println("runnable Gosched after")
	}()

	wg.Wait()

	fmt.Println("all task has done")
}

*/
