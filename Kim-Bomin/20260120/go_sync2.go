package main

import (
	"fmt"
	"runtime"
	"sync"
)

type count struct {
	num int
	// 구조체안에서 사용되는 lock/unlock 기능 호출을 위한 Mutex 필드 추가
	mutex sync.Mutex
}

func (c *count) increment() {
	c.mutex.Lock() // 임계 영역 설정
	c.num++
	c.mutex.Unlock() // 임계 영역 해제
}

func (c *count) result() {
	fmt.Println("최종 num 값:", c.num)
}

// Mutex를 사용한 동기화 예제
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
