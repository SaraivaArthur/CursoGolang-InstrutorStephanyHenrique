package main

import "fmt"

func main() {

	for numBase := 1; numBase <= 10; numBase++ {
		for multiplicado := 1; multiplicado <= 10; multiplicado++ {
			fmt.Println(numBase, "x", multiplicado, " = ", numBase*multiplicado)
		}
	}
}
