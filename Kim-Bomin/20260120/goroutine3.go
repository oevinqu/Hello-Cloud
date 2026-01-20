package main

import (
	"fmt"
	"time"
	"math/rand"
	"runtime"
)

func exe(name int) {
	r := rand.Intn(100) // 0~99
	fmt.Println(name, "start:", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>", r, i)
	}
	fmt.Println(name, "end:", time.Now())
}

func main() {

	// 1. CPU 코어 수 확인
	fmt.Println("CPU Core:", runtime.NumCPU())
	// 2. 최대 코어 수 설정 (맥스 설정)
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 3. 셋팅 됐는지 확인
	fmt.Println("After Set CPU Core:", runtime.GOMAXPROCS(0))

	fmt.Println("Go Routine Start:", time.Now())
	// 4. 고루틴 실행 (100개 동시 실행)
	for i := 0; i < 100; i++ {
		go exe(i)
	}
	// 5. 충분한 시간 대기
	time.Sleep(3 * time.Second)
	fmt.Println("Go Routine End:", time.Now())

}