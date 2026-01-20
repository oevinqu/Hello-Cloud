package main

import (
	"fmt"
)

func sum(max int) <-chan int {
	sum := 0
	c := make(chan int)
	go func() {
		for i := 1; i <= max; i++ {
			sum += i
		}
		c <- sum
	}()
	return c
}

func main() {
	c := sum(100)
	fmt.Println("1~100 합:", <-c)
}
