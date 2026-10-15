package main

import (
	"fmt"
)

type informacion struct {
	NombrePro   string
	Precio      int
	Cantidaddis int
}

func main() {

	producto1 := informacion{NombrePro: "Collar", Precio: 2, Cantidaddis: 4}
	producto2 := informacion{NombrePro: "Pulsera", Precio: 5, Cantidaddis: 10}
	producto3 := informacion{NombrePro: "Lentes", Precio: 10, Cantidaddis: 1}
	fmt.Println("La informacion del primer producto es:")
	info(producto1)
	fmt.Println("El valor total del primer producto es:")
	fmt.Println(calcular(producto1))
	fmt.Println("Y el estado stock es:")
	stock(producto1)
	// Prodcuto 2
	fmt.Println("La informacion del segundo producto es:")
	info(producto2)
	fmt.Println("El valor total del segundo prodcuto es:")
	fmt.Println(calcular(producto2))
	fmt.Println("EL estado stock del tercer prodcuto es:")
	stock(producto2)
	// tercer prodcutp
	fmt.Println("La informacion del tercer producto es:")
	info(producto3)
	fmt.Println("El valor total del tercer prodcuto es:")
	fmt.Println(calcular(producto3))
	fmt.Println("El estado stock del tercer prodcuto es:")
	stock(producto3)

}
func info(p informacion) {
	fmt.Println("Nombre:", p.NombrePro, "Precio", p.Precio, "Cantidad", p.Cantidaddis)

}
func calcular(p informacion) int {
	resultado := p.Precio * p.Cantidaddis
	return resultado
}
func stock(p informacion) {
	if p.Cantidaddis == 0 {
		fmt.Println("sin stock")
	} else {
		fmt.Println("stock")
	}

}
