package main

import (
	"fmt"
	"sync"
	"time"
)

var ss sync.WaitGroup

func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}

	close(ch)
	ss.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("读取数据%v成功\n", v)
	}
	ss.Done()
}

func main() {
	var ch1 = make(chan int, 10)

	ss.Add(1)
	go fn1(ch1)
	ss.Add(1)
	go fn2(ch1)

	ss.Wait()
}
