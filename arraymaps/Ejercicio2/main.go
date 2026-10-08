package main

import "fmt"

func main() {

	actividades := map[string]int{
		"futbol":    0,
		"basquet":   0,
		"taekwondo": 0,
		"lectura":   0,
	}

	var voto string

	fmt.Println("Actividades disponibles:")
	fmt.Println("futbol, basquet, taekwondo, lectura")

	for i := 0; i < 5; i++ {
		fmt.Print("Ingrese su voto ", i+1, ": ")
		fmt.Scan(&voto)

		actividades[voto] = actividades[voto] + 1
	}

	fmt.Println("Resultados de la votacion:")

	for actividad, cantidad := range actividades {
		fmt.Println(actividad, ":", cantidad, "votos")
	}

	determinarGanador(actividades)
}

func determinarGanador(actividades map[string]int) {

	var mayorActividad string
	var mayorVoto int = -1

	for actividad, cantidad := range actividades {

		if cantidad > mayorVoto {
			mayorVoto = cantidad
			mayorActividad = actividad
		}
	}

	fmt.Println("La actividad mas votada es:", mayorActividad, "con", mayorVoto, "votos")
}
