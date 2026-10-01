package main

import (
	"Producto/impuesto"
	"Producto/producto"
	"fmt"
)

func main() {
	var precibase float64
	fmt.Println("Ingrese su precio base")
	fmt.Scan(&precibase)
	var des float64
	fmt.Println("Ingrese el descuento:")
	fmt.Scan(&des)
	final := producto.Produc(precibase, des)
	fmt.Println("El total es:", final)
	var preciodado float64
	fmt.Println("Ingrese el precio dado para calcular su IVA:")
	fmt.Scan(&preciodado)
	var iva float64
	fmt.Println("Ingrese el IVA")
	fmt.Scan(&iva)
	preciofinal := impuesto.Impu(preciodado, iva)
	fmt.Println("El IVA total es:", preciofinal)
}
