package main

import (
	"fmt"
)

type student struct {
	FirtsName string
	LastName  string
	DNI       int
	Date      string
}

func (s student) detail() {
	fmt.Println()
}

func main() {
	stud := student{
		FirtsName: "Valentin",
		LastName:  "Faraz",
		DNI:       42062959,
		Date:      "19/12/2022",
	}
	stud.detail()
}
