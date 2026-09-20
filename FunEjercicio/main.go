package main

import "fmt"

func main() {
	var menu int

	for {
		fmt.Println("\nMenu interactivo")
		fmt.Println("1.Calcular promedio de estudiantes:")
		fmt.Println("2.Calcular la suma de los numeros ingresados")
		fmt.Println("3.Convertir la temperatura ingresada en Celsius a Fahrenheit")
		fmt.Println("4.Convertir la temperatura ingresada en Fathrenheit a Celsius")
		fmt.Println("Si desea salir del programa ingrese el numero cero")
		fmt.Println("Ingrese la opcion:")
		fmt.Scan(&menu)

		switch menu {
		case 0:
			fmt.Println("Ha salido del programa")
			return

		case 1:
			var estudiante int
			fmt.Println("Ingrese el numero de estudiante y luego las notas de cada uno: ")
			fmt.Scan(&estudiante)
			total := opcion1(estudiante)
			promedio := averageGrade(total, estudiante)
			fmt.Println("Promedio de los estudiante", promedio)
			switch {
			case promedio >= 90:
				fmt.Println("Excellent performance")
			case promedio >= 80:
				fmt.Println("Good performance")
			case promedio >= 70:
				fmt.Println("Satisfactory performance")
			default:
				fmt.Println("Needs improvement")
			}

		case 2:
			var numerosin int
			fmt.Println("Ingrese los numeros que desea sumar")
			fmt.Scan(&numerosin)
			sumatoria := opcion2(numerosin)
			fmt.Println("La sumatoria es:", sumatoria)

		case 3:
			var tempCel float64
			fmt.Println("Ingrese la temperatura en Celsius que quiera convertir a Fahrenheit")
			fmt.Scan(&tempCel)
			resultadoF := opcion3(tempCel)
			fmt.Println("La temperatura en Fahrenheit es:", resultadoF)

		case 4:
			var tempFahren float64
			fmt.Println("Ingrese la temperatura en fahrenheit que queira convertir a Celsius")
			fmt.Scan(&tempFahren)
			resultadoC := opcion4(tempFahren)
			fmt.Println("La temperatura convertida a Celsius es:", resultadoC)

		default:
			fmt.Println("Opción no válida")
		}
	}
}

func opcion1(estudiante int) float64 {
	var sumanotas float64
	var notas float64
	for i := 0; i < estudiante; i++ {
		fmt.Println("Ingresa la nota de cada estudiante")
		fmt.Scan(&notas)

		sumanotas += notas
	}
	return sumanotas
}

func averageGrade(notasingresadas float64, cantidades int) float64 {
	promedio := notasingresadas / float64(cantidades)
	if promedio >= 70 {
		fmt.Println("Aprobado")
	} else {
		fmt.Println("Reprobado")

	}
	return promedio

}

func opcion2(numerosingre int) int {
	total := 0
	for i := 1; i <= numerosingre; i++ {
		total += i
	}
	return total
}

func opcion3(tempra float64) float64 {
	conversion := (1.8 * tempra) + 32
	return conversion
}

func opcion4(temp2 float64) float64 {
	conversion2 := (temp2 - 32) / 1.8
	return conversion2

}
