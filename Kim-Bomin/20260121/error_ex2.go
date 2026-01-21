package main

import (
	"fmt"
	"math"
)

func Power(f, i float64) (float64, error) {

	if f == 0 {
		return 0, fmt.Errorf("%g은 사용 불가합니다.", f)
	}
	return math.Pow(f, i), nil // nil : 에러 없음

}

func main() {

	if f, err := Power(7, 2); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("Result 1: ", f)
	}

	if f, err := Power(0, 2); err != nil {
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("Result 2: ", f)
	}


}