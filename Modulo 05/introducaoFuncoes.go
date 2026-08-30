package main

import "fmt"

func main() {
	imprimirMensagem("mensagem x")
	imprimirMensagem("mensagem x")
}

func imprimirMensagem(mensagem string) {
	mensagem += ", bom dia"
	fmt.Println(mensagem)
}
