package main

import (
	"fmt"
)

func receiveOnly(cnt int) <-chan int {

	sum := 0
	c := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		c <- sum
		c <- 777
		c <- 777
		close(c)
	}()
	// 리턴 타입을 수신 전용 채널로 지정
	return c

}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}

func main() {

	c := receiveOnly(100)
	output := total(c)

	fmt.Println("1~100 합 + 700 + 700:", <-output)

}
