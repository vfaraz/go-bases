package main

import (
	"fmt"
	"os"
)

func main() {
	var age uint
	var employee string
	var seniority string
	var salary uint
	fmt.Println("Ingrese la edad")
	fmt.Scanf("%d", &age)
	if age < 22 {
		fmt.Println("Edad no permitida, no puede recibir prestamo")
		os.Exit(0)
	}

	fmt.Println("Tiene empleo?(si o no)")
	fmt.Scanf("%s", &employee)
	if employee == "no" {
		fmt.Println("No tiene empleo, no puede recibir prestamo")
		os.Exit(0)
	}

	fmt.Println("Tiene más de un año de antiguedad?(si o no)")
	fmt.Scanf("%s", &seniority)
	if seniority == "no" {
		fmt.Println("Necesita más de un año de antiguedad, no puede recibir prestamo")
		os.Exit(0)
	}

	fmt.Println("Ingrese el sueldo")
	fmt.Scanf("%d", &salary)
	if salary > 100000 {
		fmt.Println("Puede recibir prestamo y sin intereses debido a su sueldo")
	} else if salary <= 100000 {
		fmt.Println("Puede recibir prestamo pero con intereses")
	}
}
