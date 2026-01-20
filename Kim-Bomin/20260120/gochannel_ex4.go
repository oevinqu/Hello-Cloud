package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			// ch1에서 수신 대기
			num := <-ch1
			fmt.Println("ch1:", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for {
			// ch2로 발신
			ch2 <- "Hello Golang"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			// 여러 채널을 동시에 처리할 때 사용
			select {
			// ch1으로 발신
			case ch1 <- 777:
			// ch2에서 수신 대기
			case str := <-ch2:
				fmt.Println("ch2:", str)
			// select for문에서는 default를 잘 사용하지 않음
			// default:
			// 	fmt.Println("waiting...")
			}
		}
	}()

	time.Sleep(7 * time.Second)


}
