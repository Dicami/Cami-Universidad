package main

import "fmt"

func main() {
	var edad int = 12
	var temperatura float64 = 21.3
	var activo bool = true
	var mensaje string = "Holaaaaa"
	var dato byte = 255
	var dias [5]string = [5]string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}
	var numeros []float64 = []float64{1.1, 2.2, 3.3}
	//fmt.Printf("Edad: %v---TipoDato: %T\n", edad, edad)
	//fmt.Printf("Temperatura: %v---TipoDato: %T\n", temperatura, temperatura)
	//fmt.Printf("Activo: %v---TipoDato: %T\n", activo, activo)
	//fmt.Printf("Mensaje: %v---TipoDato: %T\n", mensaje, mensaje)
	//fmt.Printf("Dato: %v---TipoDato: %T\n", dato, dato)
	//fmt.Printf("Dias: %v---TipoDato: %T\n", dias, dias)
	//fmt.Printf("Numeros: %v---TipoDato: %T\n", numeros, numeros)
	fmt.Println("Tu edad es.", edad)
	fmt.Println("La temperatura es", temperatura)
	fmt.Println("El activo es:", activo)
	fmt.Println("El mensaje es:", mensaje)
	fmt.Println("El tipo de dato es:", dato)
	fmt.Printf("Los dias son: %v\n", dias)
	fmt.Printf("Los numeros son: %v\n", numeros)
	var i int
	var j float64
	fmt.Println("Ingresa dos valores:")
	fmt.Scanf("%d %f", &i, &j)
	fmt.Println("Resultado", (float64(i))*j*j)
	ola()
}
