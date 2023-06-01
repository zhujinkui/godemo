package main

import (
	"fmt"
	"sync"
	"time"
)

var dg sync.WaitGroup

func demo(n int) {
	for num := (n - 1) * 30000; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}

			if flag {
				fmt.Println(num, "是素数")
			}
		}
	}
	dg.Done()
}

func main() {
	start := time.Now().Unix()

	for i := 1; i <= 3; i++ {
		dg.Add(1)
		go demo(i)
	}

	dg.Wait()

	end := time.Now().Unix()

	fmt.Println(end - start)
	fmt.Println("执行完毕")
}
