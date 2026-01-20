package main

import (
	"fmt"
	"time"
)

// chan<- int: 발신 전용 채널
func sendOnly(c chan<- int, max int) {
	for i := 0; i < max; i++ {
		c <- i
		fmt.Println("send:", i)
	}
	c <- 777
	close(c)
}

// <-chan int: 수신 전용 채널
func receiveOnly(c <-chan int) {
	for v := range c {
		fmt.Println("receive:", v)
	}
	fmt.Println(<-c)
}

func main() {

	c := make(chan int)

	go sendOnly(c, 10) // 발신 전용 고루틴
	go receiveOnly(c)  // 수신 전용 고루틴

	time.Sleep(1 * time.Second)
	fmt.Println("main 종료")

}
