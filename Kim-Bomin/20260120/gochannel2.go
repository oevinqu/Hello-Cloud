package main

import (
	"fmt"
)

func rangeSum(rg int, c chan int) {
	sum := 0
	for i := 0; i < rg; i++ {
		sum += i
	}
	c <- sum
}

func main() {

	c := make(chan int)
	
	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(500, c)

	// 순서대로 값을 받음 (동기)
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
	fmt.Println("Result 3:", result3)

}