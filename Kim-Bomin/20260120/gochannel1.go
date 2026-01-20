package main

import (
	"fmt"
	"time"

)


// int형만 처리하는 채널
func work1(v chan int) {

	fmt.Println("work1 start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work1 end", time.Now())

	// 채널에 데이터 전송
	v <- 1

}

// int형만 처리하는 채널
func work2(v chan int) {

	fmt.Println("work2 start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 end", time.Now())
	// 채널에 데이터 전송
	v <- 2

}

func main() {

	fmt.Println("===Main Start===", time.Now())

	// int형만 처리하는 채널 생성
	v := make(chan int)

	go work1(v)
	go work2(v)

	// 채널에서 데이터 수신 대기 (블로킹)
	<- v
	<- v

	fmt.Println("===Main End===", time.Now())



}