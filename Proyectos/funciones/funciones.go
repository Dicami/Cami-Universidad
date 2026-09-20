package main

import "fmt"

/*
func <nombre>(param1, param2,... param n)<valores de retorno>{
	-----------------------
	-----------------------

	//return en el caso de que nuestra función retorne valores
	})

*/

func saludar() {
	fmt.Println("Hola, esta es mi primera función")
}

func bienvenida(nombre string) {
	fmt.Println("Bienvenid@", nombre)
}

func sumar_restar(num1 int, num2 int) (int, int) {
	var result_sum int
	result_sum = num1 + num2
	if num2 > num1 {
		var result_rest int
		result_rest = num2 - num1
		return result_sum, result_rest
	} else {
		fmt.Println("EL resultado es 0")
		return result_sum, 0
	}

}

//Variadic function Funciones Variadicas

func mostrarnum(numeros ...int) {
	fmt.Println("Los numeros ingresados son: ", numeros)
}

func sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	var usr string
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&usr)
	saludar()
	bienvenida(usr)

	//fmt.Println("El resultado de la suma es: ", sumar(5, 10))

	var num1, num2 int
	fmt.Println("Ingresa el primer número: ")
	fmt.Scan(&num1)
	fmt.Println("Ingresa el segundo número: ")
	fmt.Scan(&num2)
	resultadoSuma, resultadoResta := sumar_restar(num1, num2)
	fmt.Println("Los resultados son: ", resultadoSuma, resultadoResta)

	mostrarnum(5, 10, 15, 20, 25)
	fmt.Println("La sumatoria es: ", sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println("")
}
func avergeGrade(suma float64) {

}
