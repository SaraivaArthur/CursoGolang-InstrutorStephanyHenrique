package main

import (
	model "cursogolanginstrutorstephanyhenrique/model"
	"fmt"
)

func main() {
	fmt.Println("iniciando ...")

	endereco := model.Endereco{
		Rua:    "Rua x",
		Numero: 15,
		Cidade: "Osasco",
	}

	pessoa := model.Pessoa{
		Nome:     "Arthur",
		Endereco: endereco,
	}

	fmt.Println(pessoa)
	fmt.Println(endereco)

	idade := pessoa.IdadeAtual()

	fmt.Println("Idade:", idade)
	endereco.Numero = 18
}
