package main

import "fmt"

type Persona struct {
	nombre string
}

func (p Persona) GetNombre() string { //Un getter que va a retornar
	return p.nombre
}
func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre

}
func main() {
	persona := Persona{}
	persona.SetNombre("Ivan")
	fmt.Println("Su nombre es:", persona.GetNombre())
	persona.SetNombre("Yo")
	fmt.Println("Ahora su nombre es:", persona.GetNombre())

}
