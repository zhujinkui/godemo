package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()

	// 定义一个管道 10个数据的int
	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 定义一个管道 5个数据的string
	stringChan := make(chan string, 5)

	for i := 0; i < 5; i++ {
		stringChan <- "Hello" + fmt.Sprintf("%d", i)
	}

	// 使用多路复用
	for {
		select {
		case v := <-intChan:
			fmt.Println("从intChan读取数据", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Println("从stringChan读取数据", v)
			time.Sleep(time.Millisecond * 100)
		default:
			fmt.Printf("数据获取完毕")
			return
		}
	}

	end := time.Now().Unix()

	fmt.Println("进程结束", end-start)
}
