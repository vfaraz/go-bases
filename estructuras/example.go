package main

import (
	"fmt"
	//"encoding/json"
)

type Persona struct {
	Nombre   string `json:"name"`
	Apellido string `json:"last_name"`
}

func main() {
	p := Persona{"Valentin", "Faraz"}
	p.description()
}

// si coloco * en la estructura *persona (puntero)accedo a la instancia real
// si no lo coloco es una copia y no van a guardarse los cambios
func (p Persona) description() {
	fmt.Printf("%s %s\n", p.Nombre, p.Apellido)
}
