package main

import "fmt"

type listaGenerica []interface{}

func (lista *listaGenerica) remover(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[:indice], l[indice+1:]...)
	return removido
}

func (lista *listaGenerica) removerInício() interface{} {
	return lista.remover(0)
}

func (lista *listaGenerica) removerFinal() interface{} {
	return lista.remover(len(*lista) - 1)
}

func (lista *listaGenerica) removerTudo() []interface{} {
	l := *lista
	*lista = make(listaGenerica, 0, cap(*lista))
	return l
}

func main() {
	lista := listaGenerica{1, "café", 42, true, 23, "Bola", 3.14, false}

	fmt.Printf("Lista original: %v\n", lista)

	fmt.Printf("Removendo do início (%v). Depois de remover: %v\n", lista.removerInício(), lista)

	fmt.Printf("Removendo do fim (%v). Depois de remover: %v\n", lista.removerFinal(), lista)

	fmt.Printf("Removendo indice 3 (%v). Depois da remoção: %v\n", lista.remover(3), lista)

	fmt.Printf("Removendo tudo (%v). Depois da remoção: %v", lista.removerTudo(), lista)
}
