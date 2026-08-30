package main

import "fmt"

func main() {

	texto := "palavra"
	fmt.Println("Quantidade: ", len(texto)) // 'len' serve para ver quantidades
	tamanho := len(texto)
	i := 0
	for i < tamanho {
		if string(texto[i]) == "r" {
			continue // ou 'break' para parar na letra 'r'
		}
		fmt.Println(string(texto[i]))
		i++
	}
}
