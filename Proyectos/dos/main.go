package main

import "fmt"

func main() {
	num := 0
	fmt.Print("Ingresa un numero positivo: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	switch {
	case num >= 1 && num <= 10:
		fmt.Println("Número pequeño")
	case num >= 11 && num <= 100:
		fmt.Println("Número mediano")
	case num > 100:
		fmt.Println("Número grande")
	default:
		fmt.Println("Número fuera de rango o no válido")
	}
	// Ejercicio 2
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
}
