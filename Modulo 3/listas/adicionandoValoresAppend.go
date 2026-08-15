package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["sp"] = 10000000
	m["cg"] = 9000000
	m["cg"] = 7000000
	m["rj"] = 6000000
	// delete(m, "xp") // teste para deletar itens

	for chave, valor := range m {
		fmt.Println("Cidade", chave, "H:", valor)
	}
}