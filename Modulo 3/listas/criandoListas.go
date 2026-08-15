package main

import "fmt"

func main() {
	lista := []int{4, 9, 6, 7}
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista pos1: ", lista[0])
	fmt.Println("Lista pos2: ", lista[1])
	fmt.Println("Lista pos3: ", lista[2])
	fmt.Println("Tamanho da lista: ", len(lista))

	// fmt.Println("Lista pos3: ", lista[5]) // teste para ver se ele acha uma posição nao criada
}
