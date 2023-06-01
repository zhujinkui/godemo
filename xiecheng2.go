package main

import (
	"fmt"
	"sync"
	"time"
)

func test(num int) {
	defer gg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("携程[%v]打印的第%v条数据\n", num, i)
		time.Sleep(time.Millisecond * 200)
	}
}

var gg sync.WaitGroup

func main() {
	for i := 0; i < 6; i++ {
		gg.Add(1)
		go test(i)
	}

	gg.Wait()
	fmt.Println("关闭主进程")
}
