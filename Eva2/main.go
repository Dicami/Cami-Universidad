package main

import (
	"fmt"
	"time"
)

type Medicina struct {
	nombre           string
	fabricante       string
	fechaFabricacion time.Time
	vidaUtil         int
}

func NuevaMedicina(nombre, fabricante string, fecha time.Time, vida int) *Medicina {
	return &Medicina{
		nombre:           nombre,
		fabricante:       fabricante,
		fechaFabricacion: fecha,
		vidaUtil:         vida,
	}
}

func (m *Medicina) AsignarNombre(nombre string) {
	if nombre != "" {
		m.nombre = nombre
	}
}

func (m *Medicina) ObtenerNombre() string {
	return m.nombre
}

func (m *Medicina) AsignarFabricante(fabricante string) {
	if fabricante != "" {
		m.fabricante = fabricante
	}
}

func (m *Medicina) ObtenerFabricante() string {
	return m.fabricante
}

func (m *Medicina) AsignarFecha(fecha time.Time) {
	if !fecha.IsZero() {
		m.fechaFabricacion = fecha
	}
}

func (m *Medicina) ObtenerFecha() time.Time {
	return m.fechaFabricacion
}

func (m *Medicina) AsignarVidaUtil(vida int) {
	if vida > 0 {
		m.vidaUtil = vida
	}
}

func (m *Medicina) ObtenerVidaUtil() int {
	return m.vidaUtil
}

func (m *Medicina) CalcularCaducidad() time.Time {
	return m.fechaFabricacion.AddDate(0, m.vidaUtil, 0)
}

type Pastilla struct {
	Medicina
	dosisPorPastilla float64
	requiereReceta   bool
}

func NuevaPastilla(nombre, fabricante string, fecha time.Time, vida int, dosis float64, receta bool) *Pastilla {
	return &Pastilla{
		Medicina:         *NuevaMedicina(nombre, fabricante, fecha, vida),
		dosisPorPastilla: dosis,
		requiereReceta:   receta,
	}
}

func (p *Pastilla) MostrarDetalles() {
	fmt.Println("Pastilla:", p.ObtenerNombre())
	fmt.Println("Fabricante:", p.ObtenerFabricante())
	fmt.Println("Dosis:", p.dosisPorPastilla)
	fmt.Println("Requiere receta:", p.requiereReceta)
	fmt.Println("Caducidad:", p.CalcularCaducidad().Format("02/01/2006"))
	fmt.Println()
}

type Jarabe struct {
	Medicina
	volumen float64
	sabor   string
}

func NuevoJarabe(nombre, fabricante string, fecha time.Time, vida int, volumen float64, sabor string) *Jarabe {
	return &Jarabe{
		Medicina: *NuevaMedicina(nombre, fabricante, fecha, vida),
		volumen:  volumen,
		sabor:    sabor,
	}
}

func (j *Jarabe) MostrarDetalles() {
	fmt.Println("Jarabe:", j.ObtenerNombre())
	fmt.Println("Fabricante:", j.ObtenerFabricante())
	fmt.Println("Volumen:", j.volumen)
	fmt.Println("Sabor:", j.sabor)
	fmt.Println("Caducidad:", j.CalcularCaducidad().Format("02/01/2006"))
	fmt.Println()
}

func main() {
	fecha1 := time.Date(2025, 5, 10, 0, 0, 0, 0, time.UTC)
	fecha2 := time.Date(2024, 11, 15, 0, 0, 0, 0, time.UTC)

	listaPastillas := []*Pastilla{
		NuevaPastilla("Ibuprofeno", "Bayer", fecha1, 24, 500, false),
		NuevaPastilla("Amoxicilina", "Pharmabrand", fecha2, 12, 850, true),
	}

	listaJarabes := []*Jarabe{
		NuevoJarabe("Jarabe para la Tos", "Genamerica", fecha1, 18, 120.5, "Miel Limon"),
		NuevoJarabe("Ibuprofeno Pediatrico", "MK", fecha2, 6, 60.0, "Fresa"),
	}

	fmt.Println("Inventario Pastillas")
	for _, p := range listaPastillas {
		p.MostrarDetalles()
	}

	fmt.Println("Inventario Jarabes")
	for _, j := range listaJarabes {
		j.MostrarDetalles()
	}

	fmt.Println("Prueba Setter:")
	miMedicina := listaPastillas[0].Medicina
	fmt.Println("Nombre anterior:", miMedicina.ObtenerNombre())
	miMedicina.AsignarNombre("Paracetamol 1000mg")
	fmt.Println("Nombre nuevo:", miMedicina.ObtenerNombre())
}
