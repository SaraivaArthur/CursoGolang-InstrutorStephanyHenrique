package main

import (
	model "cursogolanginstrutorstephanyhenrique/model"
	"fmt"
)

func main() {
	fmt.Println("iniciando ...")

	// endereco := model.Endereco{
	// 	Rua:    "Rua x",
	// 	Numero: 15,
	// 	Cidade: "Osasco",
	// }

	// pessoa := model.Pessoa{
	// 	Nome:             "Arthur",
	// 	Endereco:         endereco,
	// 	DataDeNascimento: time.Date(2009, 02, 10, 0, 0, 0, 0, time.Local),
	// }

	// fmt.Println(pessoa)
	// fmt.Println(endereco)
	// pessoa.CalculaIdade()
	// fmt.Println(pessoa.Idade)

	automovelMoto := model.Automovel{
		Ano:    2022,
		Placa:  "XPT0",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel:   automovelMoto,
		Cilindradas: 125,
	}

	fmt.Println(moto)
	fmt.Println(moto.Modelo)
}
