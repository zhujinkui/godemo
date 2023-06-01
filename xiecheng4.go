package main

import (
	"fmt"
	"sync"
	"time"
)

var dd sync.WaitGroup

func main() {
	start := time.Now().Unix()

	ch := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		ch <- i
	}

	/*close(ch)

	for v := range ch {
		fmt.Println(v)
	}*/

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	dd.Wait()

	end := time.Now().Unix()

	fmt.Println(end - start)
	fmt.Println("执行完毕")
}
