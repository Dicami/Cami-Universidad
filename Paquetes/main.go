package main

import (
	"Paquetes/saludo"
	"Paquetes/saludo/operaciones"
	"fmt"
)

func main() {
	fmt.Println("Bienvenido a la Clase de Paquetes")

	mensaje := saludo.Saludar("Cami")
	fmt.Println(mensaje)

	suma, resta := operaciones.Sumar_restar(5, 10)
	fmt.Println("Suma:", suma, "Resta:", resta)

	total := operaciones.Sumatoria(1, 2, 3, 4)
	fmt.Println("Sumatoria:", total)
}
