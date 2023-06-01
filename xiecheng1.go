package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1(),你好 Golang - ", i)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2(),你好 Golang - ", i)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	go test1()

	wg.Add(1)

	go test2()

	for i := 0; i < 10; i++ {
		fmt.Println("main(),你好 Golang - ", i)
		time.Sleep(time.Millisecond * 100)
	}

	wg.Wait()
	fmt.Println("执行结束")
}
