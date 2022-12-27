package main

import "fmt"

func main() {
	var mes int
	meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	for {
		fmt.Println("Ingrese el numero de un mes")
		fmt.Scanf("%d", &mes)
		if mes <= 0 || mes > 12 {
			fmt.Println("El mes no existe")
		} else {
			fmt.Println("El mes es", meses[mes-1])
			break
		}
	}
}
