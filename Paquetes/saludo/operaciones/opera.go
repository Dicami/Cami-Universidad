package operaciones

import "fmt"

func Sumar_restar(num1 int, num2 int) (int, int) {
	var result_sum int
	result_sum = num1 + num2
	if num2 > num1 {
		var result_rest int
		result_rest = num2 - num1
		return result_sum, result_rest
	} else {
		fmt.Println("El resultado es 0")
		return result_sum, 0
	}
}

func Sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
