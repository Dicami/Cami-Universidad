package main

import "fmt"

type Producto struct {
	Nombre string
	precio float64
	stock  int
}

func calcularValorInvetario(p Producto) float64 {
	//multiplicar la catidad por el precio
	return p.precio * float64(p.stock)
}

func main() {
	product := Producto{}

	product.Nombre = "Intel Core Ultra 7"
	product.precio = 850.90
	product.stock = 200

	product2 := Producto{
		Nombre: "Teclado",
		precio: 30.5,
		stock:  15,
	}

	fmt.Println(product)
	fmt.Println(product2)

	fmt.Println("El Producto: ", product2.Nombre, " Tiene un valor del inventario de:", calcularValorInvetario(product2))
}
