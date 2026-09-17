package main

import "fmt"

func main() {
	//saludar()
	//nombre, edad := datos()
	//fmt.Println("Hola tu nombre es:", nombre, edad)
	var a int
	var b int
	fmt.Println("Ingrese un numero:")
	fmt.Scan(&a)
	fmt.Println("Ingrese el segundo numero:")
	fmt.Scan(&b)
	suma := sumar(a, b)
	fmt.Println("La suma es:", suma)
}

// Definir una funcion
func saludar() {
	//fmt.Println("Hola jaguar you")
	//var nombre string
	//fmt.Printf("Ingresa tu nombre")
	//fmt.Scanln(&nombre)
	//fmt.Println("Hola", nombre)
	var a int
	var b int
	fmt.Println("Ingrese un numero:")
	fmt.Scan(&a)
	fmt.Println("Ingrese el segundo numero:")
	fmt.Scan(&b)
	suma := sumar(a, b)
	fmt.Println("La suma es:", suma)

}

// funcion con retorno
func datos() (string, int) {
	var nombre string
	var edad int
	fmt.Printf("Ingresa tu nombre")
	fmt.Scanln(&nombre)
	fmt.Printf("Ingresa tu edad:")
	fmt.Scan(&edad)
	return nombre, edad
}

// Datos y argumentos
func sumar(a int, b int) int {
	suma := a + b
	return suma
}
