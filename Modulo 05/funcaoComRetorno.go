package main

import (
	"fmt"
)

func main() {
	resultado := Soma(1, 2)
	fmt.Println(resultado)
}

func Soma(numero1 int, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}
