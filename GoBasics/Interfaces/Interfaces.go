package main

import (
	"fmt"
	"time"
)

type Operação interface {
	Calcular() int
}

type Idade struct {
	anoNascimento int
}

func (i Idade) Calcular() int {
	return time.Now().Year() - i.anoNascimento
}

func (i Idade) String() string {
	return fmt.Sprintf("idade desde: %d", i.anoNascimento)
}

type Substração struct {
	Operando1 int
	Operando2 int
}

func (s Substração) Calcular() int {
	return s.Operando1 - s.Operando2
}

func (s Substração) String() string {
	return fmt.Sprintf("%d - %d", s.Operando1, s.Operando2)
}

type Soma struct {
	Operando1 int
	Operando2 int
}

func (s Soma) Calcular() int {
	return s.Operando1 + s.Operando2
}

func (s Soma) String() string {
	return fmt.Sprintf("%d + %d", s.Operando1, s.Operando2)
}

func main() {
	operacoes := make([]Operação, 4)
	operacoes[0] = Soma{10, 20}
	operacoes[1] = Substração{30, 15}
	operacoes[2] = Substração{10, 50}
	operacoes[3] = Soma{5, 2}

	fmt.Printf("Valor acumulado: %d\n", acumular(operacoes))

	idades := make([]Operação, 3)
	idades[0] = Idade{1969}
	idades[1] = Idade{1977}
	idades[2] = Idade{2001}

	fmt.Printf("Idades acumuladas: %d", acumular(idades))

}

func acumular(operacoes []Operação) int {
	acumulador := 0

	for _, v := range operacoes {
		valor := v.Calcular()
		fmt.Printf("%v = %d\n", v, valor)
		acumulador += valor
	}
	return acumulador
}
