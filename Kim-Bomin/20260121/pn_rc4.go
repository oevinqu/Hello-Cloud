package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in fileOpen:", r.(string))
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error()) // panic 발생
	} else {
		fmt.Println("File opened successfully:", f.Name()) // 실행X
	}

	defer f.Close()

}

func main() {

	fmt.Println("Starting file open operation...")

	fileOpen("nonexistent.txt") // 존재하지 않는 파일 열기 시도

	fmt.Println("Program continues after panic recovery in fileOpen.")
}