package main

import "fmt"

func main() {
	fmt.Println("aaaaaaaaaaaaaaa")
	fmt.Println("*******bucle Nomal*******")
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
	for {
		fmt.Println("Infinito")
		break
	}
	for rango := range [10]int{} {
		fmt.Println(rango)
	}
}
