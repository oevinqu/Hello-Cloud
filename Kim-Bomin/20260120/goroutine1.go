package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe 1 started", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe 1 ended", time.Now())
}

func exe2() {
	fmt.Println("exe 2 started", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe 2 ended", time.Now())
}

func exe3() {
	fmt.Println("exe 3 started", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe 3 ended", time.Now())
}

func main() {

	exe1()
	fmt.Println("Go Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Go End", time.Now())


}
