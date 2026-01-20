package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1) // CPU 코어 1개로 제한

	s := "Goroutine Closure:"

	// goroutine에서 클로저는 반복문 종료 후 실행된다
	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, ">>>>>", time.Now())
		}(i)
	}

	time.Sleep(3 * time.Second)

}
