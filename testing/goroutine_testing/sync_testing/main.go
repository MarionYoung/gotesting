package main

import (
	"fmt"
	"sync"
)
// sync.WaitGroup来实现goroutine的同步
var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i <3; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go f1(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}