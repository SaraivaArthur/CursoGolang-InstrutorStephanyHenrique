package main

import (
	"fmt"
)

func init() {
	fmt.Println("função init 1")
}

func init() {
	fmt.Println("função init 2")
}

var numeroMagico = 5

func main() {
	soma, subtracao, divisao, multiplicacao := Operacao(1, 2)
	soma -= numeroMagico
	fmt.Println(soma, subtracao, divisao, multiplicacao)
}

func Operacao(numero1 int, numero2 int) (soma int, subtracao int, divisao int, multiplicacao int) {
	numeroMagico := 3

	soma = numero1 + numero2 + numeroMagico
	subtracao = numero1 - numero2
	divisao = numero1 / numero2
	multiplicacao = numero1 * numero2
	return
}
