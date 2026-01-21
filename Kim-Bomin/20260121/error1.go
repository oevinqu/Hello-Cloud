package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("unnamedfile.txt")
	if err != nil {
		// 방법 1
		// log.Fatal(err)
		// 방법 2
		log.Fatal(err.Error())
	}
	fmt.Println(f.Name())

}