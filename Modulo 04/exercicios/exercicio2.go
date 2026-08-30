package main

import "fmt"

func main() {
	lista := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	numerosAte5 := 0
	numerosAte10 := 0
	for i := 0; i < len(lista); i++ {
		if lista[i] <= 5 {
			numerosAte5 = numerosAte5 + lista[i]
		} else {
			numerosAte10 += lista[i]
		}
	}
	fmt.Println(numerosAte5)
	fmt.Println(numerosAte10)
}