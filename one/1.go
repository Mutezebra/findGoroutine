package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	var x int
	go oneFunction(ch)
	for {
		_, _ = fmt.Scan(&x)
		ch <- x
	}
}

func oneFunction(ch chan int) {
	var x int
	for {
		fmt.Println(time.Now())
		x = <-ch
		fmt.Println(time.Now())
		fmt.Println(x)
	}
}
