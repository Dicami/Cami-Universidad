package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese una frase")
	scanner.Scan()
	texto := scanner.Text()
	palabras := strings.Fields(texto)
	veces := make(map[string]int)
	for _, palabra := range palabras {
		veces[palabra]++
	}
	fmt.Println("El mapa de palabras:", veces)
	palabraMayor := ""
	maxRepeticiones := 0
	for palabra, conteo := range veces {

		if conteo > maxRepeticiones {
			maxRepeticiones = conteo
			palabraMayor = palabra
		}
	}

	fmt.Println("La palabra que más se repite es:", palabraMayor)
}
