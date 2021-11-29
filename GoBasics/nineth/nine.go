package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go sleep(c)

	fmt.Println("Waiting...")

	end := false

	for !end {
		select {
		case end = <-c:
			fmt.Println("End!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			end = true
		}
	}
}

func sleep(c chan<- bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
