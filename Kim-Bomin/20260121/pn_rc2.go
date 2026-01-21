package main

import (
	"fmt"
)

func runFunc() {

	// panic 상태 이후 defer 함수는 실행됨
	defer func() {

		s := recover() // panic 상태 복구
		fmt.Println("Error Message:", s)

	}()

	panic("Error occured!")

	fmt.Println("End runFunc") // 실행 안됨

}

func main() {

	runFunc()

	fmt.Println("Hello Golang!") // runFunc()에서 복구되었으므로 실행됨

}
