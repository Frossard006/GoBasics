package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, v := range entrada {
		numero, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("%v é um número inválido\n", v)
			os.Exit(1)
		}
		numeros[i] = numero
	}
	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]
	n = append(n[:indicePivo], n[indicePivo+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(append(quicksort(menores), pivo), quicksort(maiores)...)
}

func particionar(n []int, p int) (menores []int, maiores []int) {
	for _, v := range n {
		if v <= p {
			menores = append(menores, v)
		} else {
			maiores = append(maiores, v)
		}
	}
	return menores, maiores
}
