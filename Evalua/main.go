package main

import "fmt"

func main ()  {

    productos:= [3] string {"Arroz", "Leche" , "Pan"}

    precio:= [3] float64 (1,25,0,95,0,50)

    valuee:= [] int {}

    nombre:= [] string {}

    subtotall:= [] int {}

    var comprar int

    fmt.Println("Ingrese el numero de productos que quiera comprar")

    fmt.Scan(&comprar)

    for i:= 0; i < comprar; i ++ {

        var nombre string

        fmt.Println("Ingresa los nombres de los productos que desea comprar")

        fmt.Scan(&nombre)

        valuee = append(valuee)

        fmt.Println("El numero de los productos son:" , valuee)



    }

    var nombreproduc string

    fmt.Println("Ingrese los nombres de los productos:")

    fmt.Scan(&nombreproduc)

    for i:= 0; i < nombreproduc; i ++ {

        nombre = append(nombre, )

    }



}
func RegistrarVenta(nombreproduc string, precio float64, cantidad float64) {
	subtotal := precio * cantidad
	
	nombre = append(nombre, nombreproduc)
	subtotall = append(subtotall, subtotal)
	
	fmt.Println("Venta registrada ")