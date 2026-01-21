package main

import (
	"fmt"
	// "log"
)

func main() {

	// Panic 함수는 호출 즉시 현재 고루틴을 패닉 상태로 만듭니다.
	// 패닉 상태가 되면 현재 함수의 실행이 즉시 중단되고,
	// 현재 함수의 실행을 중단시키는 지점까지
	// 호출된 모든 함수들이 차례대로 종료됩니다.
	// 런타임 이외에 사용자가 코드 흐름에 따라 임의로 발생 시킬 때 사용!
	// 문법적인 에러는 아니지만, 논리적인 코드 흐름에 따른 에러 상황에서 사용

	fmt.Println("Start Main")
	panic("Error occured: user stopped- panic()")
	// log.Panic("Error occured: user stopped - log.Panic()")
	fmt.Println("End Main") // 실행 안됨


}
