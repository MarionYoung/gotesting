package main

import "fmt"

func main() {
	ch := make(chan int, 1) // 缓冲区为1，每次仅能存一个值
	for i :=0; i <10; i++ {
		select {
		case x :=  <- ch: // i=1时，走此分支，i=2时走下面的分支
			fmt.Println(x)
		case ch <-i :
		}
	}
}

/*
0
2
4
6
8
*/