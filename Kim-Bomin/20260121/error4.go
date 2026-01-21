package main

import (
	"fmt"
	"log"
	"errors"
)

// 사용자 정의 에러 생성
// 리턴값은 (string, error)
func notZero(n int) (string, error) {

	if n != 0 {
		s := fmt.Sprint("This is not zero: ", n)
		return s, nil // nil : 에러 없음
	}

	return "", errors.New("0 is zero error")

}

func main() {

	// Errorf : 에러 메시지 포맷팅
	a, err := notZero(1)

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Result a: ", a)

	b, err := notZero(0)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Result b: ", b)

	fmt.Println("End Main")

}
