package main

import "fmt"

func main() {

	texto := "palavra"
	fmt.Println("Quantidade: ", len(texto)) // 'len' serve para ver quantidades
	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			continue // ou 'break' para parar na letra 'r'
		}
		fmt.Println(string(texto[i]))
	}
}