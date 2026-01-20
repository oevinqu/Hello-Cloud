package main

import (
	"fmt"
	"runtime"
)

func main() {

	// 코어 수를 1개로 제한
	runtime.GOMAXPROCS(1)
	// 채널의 버퍼 크기를 3로 설정 (3개씩 처리)
	ch := make(chan bool, 3)
	cnt := 20

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Goroutine:", i)

		}

	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main:", i)
	}

}
