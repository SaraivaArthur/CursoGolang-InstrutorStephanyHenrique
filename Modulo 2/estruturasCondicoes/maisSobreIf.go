package main

import "fmt"

func main() {
	var ehCarro bool = true
	var valorDoAutomovel = 1000.00

	if ehCarro {
		valorDoAutomovel += 55.50
	}

	fmt.Println("Valor final: ", valorDoAutomovel)
}
