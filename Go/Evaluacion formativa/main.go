package main

import (
	"fmt"
	"time"
)

// Estructura Med
type Medicina struct {
	nombre           string
	fabricante       string
	fechaFabricacion time.Time
	vidaUtil         int
}

// Constructor
func NuevaMedicina(nombre, fabricante string, fecha time.Time, vida int) *Medicina {
	return &Medicina{
		nombre:           nombre,
		fabricante:       fabricante,
		fechaFabricacion: fecha,
		vidaUtil:         vida,
	}
}

// Getters y Setters
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
	m.fechaFabricacion = fecha
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
	fecha := m.fechaFabricacion
	return fecha.AddDate(0, m.vidaUtil, 0)
}

// Estructura Pas
type Pastilla struct {
	Medicina
	dosisPorPastilla float64
	requiereReceta   bool
}

// Constructor Pastilla
func NuevaPastilla(nombre, fabricante string, fecha time.Time, vida int, dosis float64, receta bool) *Pastilla {
	return &Pastilla{
		Medicina:         *NuevaMedicina(nombre, fabricante, fecha, vida),
		dosisPorPastilla: dosis,
		requiereReceta:   receta,
	}
}

// Mostrar Pastilla
func (p *Pastilla) MostrarDetalles() {
	fmt.Println("Pastilla:", p.ObtenerNombre())
	fmt.Println("Fabricante:", p.ObtenerFabricante())
	fmt.Println("Dosis:", p.dosisPorPastilla)
	fmt.Println("Requiere receta:", p.requiereReceta)
	fmt.Println("Caducidad:", p.CalcularCaducidad().Format("02/01/2006"))
}

// Estructura Jarabe
type Jarabe struct {
	Medicina
	volumen float64
	sabor   string
}

// Constructor Jarabe
func NuevoJarabe(nombre, fabricante string, fecha time.Time, vida int, volumen float64, sabor string) *Jarabe {
	return &Jarabe{
		Medicina: *NuevaMedicina(nombre, fabricante, fecha, vida),
		volumen:  volumen,
		sabor:    sabor,
	}
}

// Mostrar Jarabe
func (j *Jarabe) MostrarDetalles() {
	fmt.Println("Jarabe:", j.ObtenerNombre())
	fmt.Println("Fabricante:", j.ObtenerFabricante())
	fmt.Println("Volumen:", j.volumen)
	fmt.Println("Sabor:", j.sabor)
	fmt.Println("Caducidad:", j.CalcularCaducidad().Format("02/01/2006"))
}

func ActualizarMedicamento(m *Medicina) {
	fmt.Println("Nombre anterior:", m.ObtenerNombre())

	m.AsignarNombre("Paracetamol 1000mg")

	fmt.Println("Nombre nuevo:", m.ObtenerNombre())
}

func main() {

	fmt.Println("Invenatio Pastilla ")

	for _, p := range listaPastillas {
		p.MostrarDetalles()

	}

	// Mostrar jarabes
	fmt.Println("Invenatio jaraabe ")

	for _, j := range listaJarabes {
		j.MostrarDetalles()
	}

}
