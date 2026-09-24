package main

import (
	contador "Practico/Contador"
	monedas "Practico/Monedas"
	"fmt"
)

func main() {
	// 1. conversor de monedas
	var dolares float64
	var moneda int

	fmt.Print("Ingrese el valor en dólares: ")
	fmt.Scan(&dolares)

	fmt.Println("Seleccione la moneda a la que desea convertir:")
	fmt.Println("1. Euros")
	fmt.Println("2. LB (Libras Esterlinas)")
	fmt.Println("3. Won (Sur Koreano)")
	fmt.Println("4. BTC")
	fmt.Print("Moneda: ")
	fmt.Scan(&moneda)

	switch moneda {
	case 1:
		fmt.Println("Valor en Euros:", monedas.ConvEuro(dolares))
	case 2:
		fmt.Println("Valor en Libras:", monedas.ConvLibra(dolares))
	case 3:
		fmt.Println("Valor en Wones:", monedas.ConvWon(dolares))
	case 4:
		fmt.Println("Valor en BTC:", monedas.ConvBTC(dolares))
	default:
		fmt.Println("Moneda no permitida.")
	}

	// 2. contador de vocales
	var frase string

	fmt.Print("\nIngrese una frase: ")
	fmt.Scan(&frase)

	a, e, i, o, u := contador.ContarVocales(frase)

	fmt.Println("Vocal a:", a)
	fmt.Println("Vocal e:", e)
	fmt.Println("Vocal i:", i)
	fmt.Println("Vocal o:", o)
	fmt.Println("Vocal u:", u)
}
