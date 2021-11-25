package main

import (
	"errors"
	"fmt"
)

type Pilha struct {
	valores []interface{}
}

func (p Pilha) Tamanho() int {
	return len(p.valores)
}

func (p Pilha) Vazia() bool {
	return p.Tamanho() == 0
}

func (p *Pilha) Empilhar(valor ...interface{}) {
	for i := 0; i < len(valor); i++ {
		p.valores = append(p.valores, valor[i])
	}

}

func (p *Pilha) Desempilhar() (interface{}, error) {
	if p.Vazia() {
		return nil, errors.New("pilha já está vazia")
	}
	valor := p.valores[p.Tamanho()-1]
	p.valores = p.valores[:p.Tamanho()-1]
	return valor, nil
}

func main() {
	pilha := Pilha{}
	fmt.Println("pilha criada com tamanho:", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())

	pilha.Empilhar("FIM", 47, "yebgeg", true, "eu")
	fmt.Println("valor após empilhar:", pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())

	for !pilha.Vazia() {
		v, err := pilha.Desempilhar()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("desempilhando:", v)
		fmt.Println("tamanho:", pilha.Tamanho())
		fmt.Println("Vazia?", pilha.Vazia())
	}
}
