package main

import "fmt"

func main() {
	var ehCarro bool = false
	var valorDoAutomovel = 1000.00

	if ehCarro {
		valorDoAutomovel += 55.50
	}

	if !ehCarro {
		valorDoAutomovel += 20.50
	}

	fmt.Println("Valor final: ", valorDoAutomovel)
}
