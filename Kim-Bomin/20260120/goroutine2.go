package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println("exe start", time.Now())

	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>", i)
	}

	fmt.Println("exe end", time.Now())
}

func main() {

	exe("t1")

	fmt.Println("go routine start :", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(5 * time.Second)
	fmt.Println("go routine end :", time.Now())

}
