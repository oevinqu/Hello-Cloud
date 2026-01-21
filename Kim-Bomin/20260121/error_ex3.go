package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 예외 처리 커스텀 구조체
type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 에러 발생 값
	message string    // 에러 메시지

}

// 예외 처리 커스텀 구조체의 Error() 메서드 구현
func (e PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {

	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용 불가합니다."}
	}

	// 숫자 이외의 값 체크 (NaN: Not a Number)
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "f가 숫자가 아닙니다."}
	}

	// 숫자 이외의 값 체크 (NaN: Not a Number)
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: i, message: "i가 숫자가 아닙니다."}
	}

	return math.Pow(f, i), nil // nil : 에러 없음

}

func main() {

	// f가 0일 때
	v, err := Power(0, 2)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Result 1: ", v)
	}

	// f가 NaN일 때
	v, err = Power(math.NaN(), 2)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Result 2: ", v)
	}

	// i가 NaN일 때
	v, err = Power(7, math.NaN())
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Result 3: ", v)
	}

	// 정상 케이스
	v, err = Power(7, 2)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println("Result 4: ", v)
	}

}
