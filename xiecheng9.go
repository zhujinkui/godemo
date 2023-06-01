package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

var hh sync.WaitGroup
var mutex sync.Mutex

func test5() {
	mutex.Lock()
	count++
	fmt.Println("the count is :", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	hh.Done()
}

func main() {
	start := time.Now().Unix()

	for i := 0; i < 40; i++ {
		hh.Add(1)
		go test5()
	}

	hh.Wait()

	end := time.Now().Unix()

	fmt.Println("进程结束", end-start)
}
