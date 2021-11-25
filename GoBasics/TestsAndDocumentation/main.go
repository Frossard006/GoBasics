package main

import (
	quote "Seventh/foo"
	"Seventh/sum"
	"fmt"
)

func main() {
	fmt.Println(sum.Count(quote.SunAlso))

	for k, v := range sum.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
