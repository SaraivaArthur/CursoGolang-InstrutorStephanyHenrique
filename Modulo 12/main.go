package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")
		}
	}()

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("FileNotExist")
	}
}

func main() {

	// parte 2

	ReadFile()
	fmt.Println("Fim.")

	// parte 1

	// file, err := os.Create("./settings.txt")
	// defer file.Close()
	// defer ShowText()

	// if err != nil {
	// 	panic(err)
	// }

	// _, err = file.Write([]byte("teste"))
}

// func ShowText() {
// 	fmt.Println("Finalizando de manipular o arquivo.")
// }
