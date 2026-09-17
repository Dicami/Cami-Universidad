package main

import "fmt"

func main() {
	numero := 10
	puntero := &numero
	fmt.Println("Valor:", numero)
	fmt.Println("Direccion de memoria:", &numero) //contiene la direccion de memoria de numero
	fmt.Println("Puntero", puntero)
	fmt.Println("Puntero", *puntero) // valor de puntero
}
