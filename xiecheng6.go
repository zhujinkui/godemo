package main

import (
	"fmt"
	"sync"
	"time"
)

var ee sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i <= 500000; i++ {
		intChan <- i
	}

	close(intChan)
	ee.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	// close(primeChan)
	ee.Done()

	exitChan <- true
}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	ee.Done()
}

func main() {
	start := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 10)

	// 存放携程
	ee.Add(1)
	go putNum(intChan)

	// 统计携程
	for i := 1; i <= 10; i++ {
		ee.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 打印携程
	ee.Add(1)
	go printPrime(primeChan)

	// exitChan 是否存满值
	ee.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			<-exitChan
		}

		close(primeChan)
		ee.Done()
	}()

	ee.Wait()
	end := time.Now().Unix()

	fmt.Println("进程结束", end-start)
}
