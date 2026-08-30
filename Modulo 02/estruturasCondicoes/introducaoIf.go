package main

import "fmt"

func main() {
	salario := 850.00
	var salarioMaisOBonus float64

	salarioMaisOBonus = salario

	if salario < 1000 {
		salarioMaisOBonus = (salarioMaisOBonus + 100)
	}

	fmt.Println("Salário: ", salarioMaisOBonus)
}
