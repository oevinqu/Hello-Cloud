package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num++
}

func (c *count) result() {
	fmt.Println("최종 num 값:", c.num)
}

// 경쟁 상태 예제
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// 구조체 선언
	c := count{num: 0}

	done := make(chan bool)

	for i := 0; i < 10000; i++ {
		go func() {
			c.increment()
			done <- true
			// 다른 CPU에 양보
			runtime.Gosched()
		}()
	}

	for i := 0; i < 10000; i++ {
		<-done
	}

	c.result()

}
