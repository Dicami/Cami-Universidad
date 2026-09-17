package main

import "fmt"

func main() {
	var n int
	fmt.Print("Ingresa un número: ")
	fmt.Scan(&n)

	
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}

	
	if n%2 == 0 {
		fmt.Println("Es Par")
	} else {
		fmt.Println("Es Impar")
	}

	
	switch {
	case n <= 5:
		fmt.Println("Número pequeño")
	case n <= 10:
		fmt.Println("Número mediano")
	default:
		fmt.Println("Número grande")
	}
	// Ejercicio 2 
}