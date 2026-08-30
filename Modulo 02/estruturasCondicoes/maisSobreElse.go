package main

import "fmt"

func main() {
	salario := 1000.00

	if salario < 1000.00 {
		salario -= (salario * 0.08)
	} else if salario <= 1200.00 {
		salario -= (salario * 0.10)
	} else {
		salario -= (salario * 0.15)
	}

	fmt.Println("Salário final: ", salario)
}
