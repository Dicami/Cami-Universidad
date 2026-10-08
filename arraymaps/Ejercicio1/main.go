package main

import "fmt"

func main() {

	var notas [6][4]float64

	for i := 0; i < 6; i++ {
		fmt.Println("Estudiante", i+1)

		for j := 0; j < 4; j++ {
			fmt.Print("Ingrese la nota de la materia ", j+1, ": ")
			fmt.Scan(&notas[i][j])
		}
	}

	var sumaTotal float64

	for i := 0; i < 6; i++ {

		var estudiante []float64 = notas[i][:]
		var suma float64 = 0
		var max float64 = estudiante[0]
		var min float64 = estudiante[0]

		for j := 0; j < 4; j++ {
			var nota float64 = estudiante[j]

			suma = suma + nota

			if nota > max {
				max = nota
			}

			if nota < min {
				min = nota
			}
		}

		var promedio float64 = suma / 4
		sumaTotal = sumaTotal + suma

		fmt.Println("Estudiante", i+1)
		fmt.Println("Promedio:", promedio)
		fmt.Println("Nota más alta:", max)
		fmt.Println("Nota más baja:", min)
	}

	var promedioGeneral float64 = sumaTotal / 24

	fmt.Println("Promedio general de la clase:", promedioGeneral)
}
