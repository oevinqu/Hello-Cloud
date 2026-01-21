package main

import (
	"fmt"
	"errors"
)

func main() {

	var err1 error = errors.New("This is a error 1")
	err2 := fmt.Errorf("This is a error 2")

	fmt.Println(err1)
	fmt.Println(err2.Error())

}