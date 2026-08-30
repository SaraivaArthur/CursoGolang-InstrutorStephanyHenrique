package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 9000000
	m["cg"] = 7000000

	valor, foiEncontrado := m["rj"]
	if foiEncontrado {
		fmt.Println(valor)
	} else {
		fmt.Println("Chave não existe")
	}

	fmt.Println(m)
}
