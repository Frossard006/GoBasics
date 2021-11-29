package main

import (
	"fmt"
)

func main() {
	separa(1, 5, 4, 4, 7, 8, 9, 1, 7, 8, 6, 8)

}

func separa(x ...int) {
	par := make(chan int)
	impar := make(chan int)
	pronto := make(chan int)
	o := 1
	go func() {
		for _, v := range x {
			if v%2 == 0 {
				par <- v
			} else {
				impar <- v
			}
		}
		pronto <- 0
		close(par)
		close(impar)
	}()

	for {
		select {
		case v1 := <-par:
			fmt.Println("par", v1)
		case v2 := <-impar:
			fmt.Println("impar", v2)
		case o = <-pronto:
			fmt.Println(o)
			return
		}
	}
}
