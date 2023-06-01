package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello word")
	}
}

func test10() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang"
}

func main() {
	start := time.Now().Unix()

	go sayHello()
	go test10()

	time.Sleep(time.Second)

	end := time.Now().Unix()

	fmt.Println("进程结束", end-start)
}
