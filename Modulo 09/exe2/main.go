package main

import (
	"exe1/model"
	"fmt"
	"time"
)

func main() {

	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "Arroz")
	nomeDosItens = append(nomeDosItens, "Feijão")
	nomeDosItens = append(nomeDosItens, "Carne")
	nomeDosItens = append(nomeDosItens, "Sabonete")

	compra, err := model.NewCompra("Mercado do Zé", time.Now(), nomeDosItens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}
}
