package test1

import (
	"fmt"
	"sync"
	"time"
)

// 无缓冲通道 读写都是阻塞的
var ch = make(chan int)
var wg1 sync.WaitGroup

func r_ch(ch chan int) {
	defer wg1.Done()

	value := <-ch

	fmt.Printf("read value: %d\n", value)
}

func w_ch(ch chan int) {
	defer wg1.Done()

	time.Sleep(2 * time.Second)

	ch <- 100
	fmt.Printf("write value \n")
}

func main() {

	wg1.Add(2)

	go r_ch(ch)
	go w_ch(ch)

	wg1.Wait()
	fmt.Println("all task has done")
}
