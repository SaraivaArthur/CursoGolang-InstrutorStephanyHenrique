package main

import (
	"fmt"
)

// tem que importar biblioteca "math" e "errors"

// parte 1

// type geometria interface {
// 	area() float64
// }

// type retangulo struct {
// 	largura, altura float64
// }

// type circulo struct {
// 	radius float64
// }

// func (r retangulo) area() float64 {
// 	return r.largura * r.altura
// }

// func (c circulo) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func ExibiGeometria(g geometria) {
// 	fmt.Println(g.area())
// }

// retangulo := retangulo{
// 	largura: 1,
// 	altura:  2,
// }

// circulo := circulo{
// 	radius: 3,
// }

// ExibiGeometria(retangulo)
// ExibiGeometria(circulo)

// parte 2

// ExibiError(errors.New("a error"))

// p := ProblemaDeNetwork{
// 	rede:     true,
// 	hardware: false,
// }
// ExibiError(errors.New("a error"))
// ExibiError(p)
// }

// type ProblemaDeNetwork struct {
// 	rede     bool
// 	hardware bool
// }

// func (p ProblemaDeNetwork) Error() string {
// 	if p.rede {
// 		return "Problema de rede"
// 	} else if p.hardware {
// 		return "Problema de hardware"
// 	} else {
// 		return "outro problema"
// 	}
// }

// func ExibiError(err error) {
// 	fmt.Println(err)
// }

func main() {
	fmt.Println("inicializando ...")

	// parte 3

	var lista []interface{}
	lista = append(lista, 10)
	lista = append(lista, 7.5)
	lista = append(lista, true)
	lista = append(lista, "teste")

	for _, valor := range lista {

		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")
		} else {
			fmt.Println(valor)
		}
	}
}
