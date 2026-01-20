package main

import (
	"fmt"
)

func main() {

	ch := make(chan bool)
	go func() {
		for i := 0 ; i < 5 ; i++ {
			ch <- true
			fmt.Println("Goroutine:", i)
		}
		close(ch)
	}()

	// 채널이 닫힐 때까지 반복
	for i := range ch {
		fmt.Println("ex:", i)
	}


}