package main

import "fmt"

var nombre []string
var subtotall []float64

func RegistrarVenta(nombreproduc string, precio float64, cantidad float64) {
	subtotal := precio * cantidad

	nombre = append(nombre, nombreproduc)
	subtotall = append(subtotall, subtotal)

	fmt.Println("Venta registrada ")
}

func MostrarEstadisticas() {
	if len(subtotall) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	var total float64 = 0
	for i := 0; i < len(subtotall); i++ {
		total = total + subtotall[i]
	}

	fmt.Println("El total recaudado es:", total)
}

func main() {
	productos := [3]string{"Arroz", "Leche", "Pan"}
	precio := [3]float64{1.25, 0.95, 0.50}

	var opcion int = 0

	for opcion != 3 {
		fmt.Println("MENU ")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadisticas")
		fmt.Println("3. Salir")
		fmt.Println("Seleccione una opcion:")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Lista de productos:")
			fmt.Println("1.", productos[0], "-", precio[0])
			fmt.Println("2.", productos[1], "-", precio[1])
			fmt.Println("3.", productos[2], "-", precio[2])

			var comprar int
			fmt.Println("Seleccione el numero de producto (1, 2 o 3):")
			fmt.Scan(&comprar)

			var cantidad float64
			fmt.Println("Ingrese la cantidad vendida:")
			fmt.Scan(&cantidad)

			switch comprar {
			case 1:
				RegistrarVenta(productos[0], precio[0], cantidad)
			case 2:
				RegistrarVenta(productos[1], precio[1], cantidad)
			case 3:
				RegistrarVenta(productos[2], precio[2], cantidad)
			default:
				fmt.Println("Producto invalido.")
			}

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("Saliendo del programa...")

		default:
			fmt.Println("Opcion invalida.")
		}
	}
}
