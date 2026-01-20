package main

import (
	"fmt"
)

func main() {

	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Goroutine:", i)

		}

	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main:", i)
	}

}
